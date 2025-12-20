```go
// Copyright (c) 2024 - The fintech-bridge-platform Authors
// All rights reserved.
//
// You may use, distribute and modify this code under the
// terms of the Apache License, Version 2.0.
//
// You should have received a copy of the Apache License
// along with this program. If not, see <http://www.apache.org/licenses/LICENSE-2.0>.

package risk

import (
	"context"
	"log"

	"github.com/fintech-bridge/platform/pkg/commands"
	"github.com/fintech-bridge/platform/pkg/events"
)

// Rule defines the interface for a single, stateless risk evaluation rule.
// Each rule inspects an event and, if its conditions are met, generates a
// corresponding command to mitigate the identified risk. This interface-based
// approach allows for modular and testable risk logic.
type Rule interface {
	// Name returns a unique, human-readable identifier for the rule.
	// This is used for logging and monitoring purposes.
	Name() string

	// Evaluate processes a given event against the rule's logic.
	// It returns a command if the rule is triggered, a boolean indicating
	// whether it was triggered, and an error if the evaluation fails.
	// If the rule is not triggered, the returned command should be nil.
	// The evaluation must be deterministic and should not have side effects.
	Evaluate(ctx context.Context, event events.Event) (commands.Command, bool, error)
}

// FeedbackLoop is a critical component that generates new commands in response
// to risk events, creating an automated, real-time control loop. It consumes
// events from a source, evaluates them against a set of configurable rules,
// and dispatches corrective commands to a sink.
type FeedbackLoop struct {
	eventSource  <-chan events.Event
	commandSink  chan<- commands.Command
	rules        []Rule
	logger       *log.Logger
	stopChan     chan struct{}
	shutdownChan chan struct{}
}

// NewFeedbackLoop creates and initializes a new FeedbackLoop.
// It requires a source for events, a sink for commands, a slice of rules to
// apply, and a logger for auditable output.
func NewFeedbackLoop(
	eventSource <-chan events.Event,
	commandSink chan<- commands.Command,
	rules []Rule,
	logger *log.Logger,
) *FeedbackLoop {
	if eventSource == nil || commandSink == nil || logger == nil {
		// A nil dependency would cause a panic. Fail fast.
		panic("risk.NewFeedbackLoop: received nil dependency")
	}
	return &FeedbackLoop{
		eventSource:  eventSource,
		commandSink:  commandSink,
		rules:        rules,
		logger:       logger,
		stopChan:     make(chan struct{}),
		shutdownChan: make(chan struct{}),
	}
}

// Start launches the feedback loop's main processing goroutine.
// It will begin consuming and processing events immediately. The provided
// context can be used to signal a shutdown request.
func (fl *FeedbackLoop) Start(ctx context.Context) {
	fl.logger.Println("Starting risk feedback loop...")
	go fl.run(ctx)
}

// Stop signals the feedback loop to shut down gracefully.
// It blocks until the main processing goroutine has terminated.
func (fl *FeedbackLoop) Stop() {
	fl.logger.Println("Stopping risk feedback loop...")
	close(fl.stopChan)
	<-fl.shutdownChan
	fl.logger.Println("Risk feedback loop stopped.")
}

// run is the core event processing loop. It listens for events on the
// eventSource channel and dispatches them for rule evaluation. It respects
// cancellation signals from the parent context and the Stop() method.
func (fl *FeedbackLoop) run(ctx context.Context) {
	defer close(fl.shutdownChan)

	for {
		select {
		case <-fl.stopChan:
			return
		case <-ctx.Done():
			fl.logger.Printf("Context cancelled, shutting down feedback loop: %v", ctx.Err())
			return
		case event, ok := <-fl.eventSource:
			if !ok {
				fl.logger.Println("Event source channel closed, shutting down.")
				return
			}
			fl.processEvent(ctx, event)
		}
	}
}

// processEvent iterates through all registered rules and evaluates the given
// event against each one. If a rule is triggered, its generated command is
// dispatched to the command sink.
func (fl *FeedbackLoop) processEvent(ctx context.Context, event events.Event) {
	fl.logger.Printf("Processing event: type=%T, id=%s", event, event.ID())

	for _, rule := range fl.rules {
		cmd, triggered, err := rule.Evaluate(ctx, event)
		if err != nil {
			// Fail-closed semantics: An error during rule evaluation is a
			// critical failure. It indicates a bug or a state of uncertainty
			// that could compromise system safety. We log it as a critical
			// error. A more aggressive implementation might dispatch a
			// system-wide halt command here.
			fl.logger.Printf(
				"CRITICAL: Rule evaluation failed: rule=%s, event_id=%s, error=%v",
				rule.Name(), event.ID(), err,
			)
			// Continue to the next rule to ensure other risk checks are still performed.
			continue
		}

		if triggered {
			fl.logger.Printf(
				"Rule triggered: rule=%s, event_id=%s, generated_command=%T",
				rule.Name(), event.ID(), cmd,
			)
			fl.dispatchCommand(ctx, cmd)
			// Note: We allow multiple rules to trigger on a single event.
			// This enables layered risk responses (e.g., flag for review AND
			// reduce trading limits). Rule design should account for this.
		}
	}
}

// dispatchCommand sends a command to the command sink, respecting context
// cancellation and shutdown signals.
func (fl *FeedbackLoop) dispatchCommand(ctx context.Context, cmd commands.Command) {
	if cmd == nil {
		// This can happen if a rule triggers but decides not to issue a command,
		// which is a valid but unusual case.
		return
	}

	select {
	case fl.commandSink <- cmd:
		fl.logger.Printf("Dispatched command: type=%T, id=%s, target_id=%s", cmd, cmd.ID(), cmd.TargetID())
	case <-ctx.Done():
		fl.logger.Printf("Failed to dispatch command %T (%s): context cancelled.", cmd, cmd.ID())
	case <-fl.stopChan:
		fl.logger.Printf("Failed to dispatch command %T (%s): feedback loop shutting down.", cmd, cmd.ID())
	}
}
### END_OF_FILE_COMPLETED ###
```