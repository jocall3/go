```go
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/google/wire"

	"github.com/Tender-Services/bridge-go/internal/api"
	"github.com/Tender-Services/bridge-go/internal/api/handler"
	"github.com/Tender-Services/bridge-go/internal/config"
	"github.com/Tender-Services/bridge-go/internal/domain/account"
	"github.com/Tender-Services/bridge-go/internal/domain/instrument"
	"github.com/Tender-Services/bridge-go/internal/domain/ledger"
	"github.com/Tender-Services/bridge-go/internal/domain/transaction"
	"github.com/Tender-Services/bridge-go/internal/infrastructure/database"
	"github.com/Tender-Services/bridge-go/internal/infrastructure/logging"
	"github.com/Tender-Services/bridge-go/internal/infrastructure/repositories/accountrepo"
	"github.com/Tender-Services/bridge-go/internal/infrastructure/repositories/instrumentrepo"
	"github.com/Tender-Services/bridge-go/internal/infrastructure/repositories/ledgerrepo"
	"github.com/Tender-Services/bridge-go/internal/infrastructure/repositories/transactionrepo"
	"github.com/Tender-Services/bridge-go/internal/server"
	"github.com/Tender-Services/bridge-go/internal/service/accountsvc"
	"github.com/Tender-Services/bridge-go/internal/service/instrumentsvc"
	"github.com/Tender-Services/bridge-go/internal/service/ledgersvc"
	"github.com/Tender-Services/bridge-go/internal/service/risksvc"
	"github.com/Tender-Services/bridge-go/internal/service/settlementsvc"
	"github.com/Tender-Services/bridge-go/internal/service/transactionsvc"
)

// appSet provides all the top-level application components.
var appSet = wire.NewSet(
	config.New,
	logging.NewLogger,
	database.NewDB,
	server.New,
)

// repositorySet provides all the repository implementations and binds them to their interfaces.
var repositorySet = wire.NewSet(
	accountrepo.New,
	wire.Bind(new(account.Repository), new(*accountrepo.PostgresRepository)),

	transactionrepo.New,
	wire.Bind(new(transaction.Repository), new(*transactionrepo.PostgresRepository)),

	ledgerrepo.New,
	wire.Bind(new(ledger.Repository), new(*ledgerrepo.PostgresRepository)),

	instrumentrepo.New,
	wire.Bind(new(instrument.Repository), new(*instrumentrepo.PostgresRepository)),
)

// serviceSet provides all the business logic services.
var serviceSet = wire.NewSet(
	accountsvc.New,
	transactionsvc.New,
	ledgersvc.New,
	risksvc.NewEngine,
	settlementsvc.NewEngine,
	instrumentsvc.New,
)

// apiSet provides the API router and all its handlers.
var apiSet = wire.NewSet(
	handler.NewAccountHandler,
	handler.NewTransactionHandler,
	handler.NewHealthHandler,
	handler.NewInstrumentHandler,
	api.NewRouter,
)

// InitializeServer creates a new server with all its dependencies.
// The cleanup function should be deferred by the caller to ensure resources are released.
func InitializeServer(ctx context.Context) (*server.Server, func(), error) {
	wire.Build(
		appSet,
		repositorySet,
		serviceSet,
		apiSet,
	)
	return nil, nil, nil // This will be replaced by Wire's generated code
}

```