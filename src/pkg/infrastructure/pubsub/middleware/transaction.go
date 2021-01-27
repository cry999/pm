package middleware

// TODO: persistence に移すべき?

import (
	"context"
	"io"

	"github.com/cry999/pm-projects/pkg/infrastructure/persistence"
	"github.com/cry999/pm-projects/pkg/interfaces/logger"
	"github.com/cry999/pm-projects/pkg/interfaces/pubsub"
)

var transactionLogger = logger.NewDefaultLogger(logger.LoggerLevelDebug, "pubsub")

// Logger ...
func Logger() logger.Logger {
	return transactionLogger
}

// Transaction ...
func Transaction(next pubsub.Subscriber) pubsub.Subscriber {
	return func(ctx context.Context, payload io.Reader) (err error) {
		db, err := persistence.OpenMySQL()
		if err != nil {
			Logger().Error("failed to connect mysql: %v", err)
			return
		}
		defer db.Close()

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			Logger().Error("failed to begin transaction: %v", err)
			return
		}
		defer func() {
			if err != nil {
				Logger().Error("rollback because: %v", err)
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
		persistence.SetTransaction(ctx, tx)
		return next(ctx, payload)
	}
}
