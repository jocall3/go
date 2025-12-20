```go
// Package settlement provides the components responsible for the final,
// irreversible state changes in the system. It ensures that all financial
// operations are processed in a deterministic and auditable manner.
package settlement

import (
	"context"
	"fmt"
	"sync"
	"time"

	// In a real project, the event package would define the core event types.
	// For example: "github.com/your-org/your-project/pkg/event"
	"github.com/your-org/financial-infrastructure/pkg/event"
)

// Sequencer ensures that events are processed in the exact order of their sequence IDs.
// It acts as a gatekeeper for a Projector, guaranteeing that state is reconstructed
// deterministically by ordering a potentially concurrent and out-of-order stream of events.
//
// The Sequencer holds out-of-order events in a pending buffer until the missing
// events arrive. This prevents race conditions and ensures that state transitions
// are applied correctly and sequentially, embodying the "fail-closed" and
// "deterministic behavior" principles of the system.
type Sequencer struct {
	// expectedSequence is the next sequence ID the sequencer is waiting for.
	expectedSequence uint64

	// pendingEvents holds events that have arrived out of order.
	// They are keyed by their sequence ID.
	// NOTE: In a system with potential for large, sustained gaps in the event
	// stream, this buffer could grow indefinitely. Production monitoring should
	// track its size (`PendingCount`) to detect such scenarios.
	pendingEvents map[uint64]event.Sequenced

	// inputCh is the channel for receiving unordered events from various sources.
	inputCh chan event.Sequenced

	// outputCh is the channel where ordered events are sent to the consumer (e.g., a Projector).
	outputCh chan event.Sequenced

	// mu protects access to internal state (expectedSequence, pendingEvents).
	mu sync.Mutex

	// wg is used to wait for the processing goroutine to finish during shutdown.
	wg sync.WaitGroup

	// cancel is used to signal the processing goroutine to stop.
	cancel context.CancelFunc
}

// NewSequencer creates and initializes a new Sequencer.
// startSequence should be the sequence ID of the *next* event to be processed.
// For a new system, this is typically 1. For a recovering system, it would be
// the last successfully processed sequence ID + 1, retrieved from a persistent store.
// bufferSize defines the capacity of the input and output channels.
func NewSequencer(startSequence uint64, bufferSize int) *Sequencer {
	if startSequence == 0 {
		// Sequence IDs are 1-based to avoid ambiguity with a zero value.
		startSequence = 1
	}
	return &Sequencer{
		expectedSequence: startSequence,
		pendingEvents:    make(map[uint64]event.Sequenced),
		inputCh:          make(chan event.Sequenced, bufferSize),
		outputCh:         make(chan event.Sequenced, bufferSize),
	}
}

// Start begins the sequencer's event processing loop in a new goroutine.
// It requires a parent context for managing its lifecycle.
func (s *Sequencer) Start(ctx context.Context) {
	ctx, s.cancel = context.WithCancel(ctx)
	s.wg.Add(1)
	go s.run(ctx)
}

// Stop gracefully shuts down the sequencer.
// It waits for the processing loop to finish handling any in-flight events
// before closing the output channel.
func (s *Sequencer) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
	s.wg.Wait()
}

// Submit adds an event to the sequencer's processing queue.
// This method is safe for concurrent use by multiple producers.
// It returns an error if the sequencer's context has been canceled, preventing
// submission to a stopped sequencer.
func (s *Sequencer) Submit(ctx context.Context, ev event.Sequenced) error {
	select {
	case s.inputCh <- ev:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("failed to submit event: sequencer context is done: %w", ctx.Err())
	}
}

// Ordered returns a read-only channel of sequenced events.
// Consumers, such as a state projector, should read from this channel to receive events
// in their correct, deterministic order.
func (s *Sequencer) Ordered() <-chan event.Sequenced {
	return s.outputCh
}

// run is the core processing loop of the sequencer.
// It reads from the input channel, orders events, and sends them to the output channel.
// This method should not be called directly; use Start().
func (s *Sequencer) run(ctx context.Context) {
	defer s.wg.Done()
	defer close(s.outputCh) // Ensure output is closed when loop exits.

	for {
		select {
		case ev, ok := <-s.inputCh:
			if !ok {
				// inputCh was closed, which only happens during shutdown.
				// The loop will naturally exit after this.
				return
			}
			s.processEvent(ctx, ev)
		case <-ctx.Done():
			// Context was canceled, time to shut down.
			// We no longer accept new events via Submit, but we
			// drain the input channel to process any buffered events
			// before exiting.
			s.drainAndProcess(ctx)
			return
		}
	}
}

// drainAndProcess handles any remaining events in the input channel during shutdown.
func (s *Sequencer) drainAndProcess(ctx context.Context) {
	close(s.inputCh) // Stop accepting new events.
	for ev := range s.inputCh {
		s.processEvent(ctx, ev)
	}
}

// processEvent is the heart of the sequencing logic. It is called for each
// incoming event.
func (s *Sequencer) processEvent(ctx context.Context, ev event.Sequenced) {
	s.mu.Lock()
	defer s.mu.Unlock()

	seqID := ev.SequenceID()

	// Discard events that have already been processed.
	// This can happen during recovery or with at-least-once delivery systems.
	if seqID < s.expectedSequence {
		// TODO: Add structured logging here to monitor for duplicate events.
		// e.g., log.Warn("sequencer: discarded duplicate event", "sequenceID", seqID, "expected", s.expectedSequence)
		return
	}

	// If the event is the one we are waiting for, process it.
	if seqID == s.expectedSequence {
		s.pushToOutput(ctx, ev)
		s.expectedSequence++
		s.processPending(ctx) // Check if buffered events can now be processed.
	} else {
		// Event arrived out of order, buffer it for later.
		// This prevents overwriting if the same out-of-order event is sent twice.
		if _, exists := s.pendingEvents[seqID]; !exists {
			s.pendingEvents[seqID] = ev
		}
	}
}

// processPending checks the buffer for the next expected event and processes
// all consecutive events that are now available.
// This method must be called with the mutex held.
func (s *Sequencer) processPending(ctx context.Context) {
	for {
		// Check if the context has been cancelled before processing the next pending item.
		if ctx.Err() != nil {
			return
		}

		nextEv, found := s.pendingEvents[s.expectedSequence]
		if !found {
			// The next required event is not in the buffer, so we stop.
			break
		}

		// We found the next event in the sequence.
		s.pushToOutput(ctx, nextEv)
		delete(s.pendingEvents, s.expectedSequence)
		s.expectedSequence++
	}
}

// pushToOutput sends a sequenced event to the output channel.
// It handles the case where the output channel might be blocked and the
// context is canceled, preventing deadlocks during shutdown.
// This method must be called with the mutex held.
func (s *Sequencer) pushToOutput(ctx context.Context, ev event.Sequenced) {
	// Blocking here is a feature, not a bug. It provides backpressure.
	// If the consumer (projector) is slow or stuck, the entire pipeline
	// pauses, preventing state divergence or unbounded memory growth.
	// The timeout ensures that a completely dead consumer doesn't lock
	// the system forever, adhering to fail-closed semantics.
	select {
	case s.outputCh <- ev:
		// Event successfully sent.
	case <-ctx.Done():
		// The sequencer is shutting down, so we don't send the event.
		// The event remains in the pending map and will be re-processed on restart.
	case <-time.After(30 * time.Second):
		// This is a critical failure. The consumer (projector) is stuck.
		// Halting is the safest option. A supervisor process should detect the
		// panic and take corrective action (e.g., restart, alert).
		panic(fmt.Sprintf(
			"sequencer: consumer failed to read from output channel for 30s. Halting. Last processed sequence: %d",
			s.expectedSequence-1,
		))
	}
}

// PendingCount returns the number of events currently held in the pending buffer.
// Useful for monitoring and diagnostics.
func (s *Sequencer) PendingCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.pendingEvents)
}

// ExpectedSequence returns the next sequence ID the sequencer is waiting for.
// Useful for monitoring and diagnostics.
func (s *Sequencer) ExpectedSequence() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.expectedSequence
}

```