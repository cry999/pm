package middleware

import (
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
)

// Transaction ...
func Transaction(next web.HandlerFunc) web.HandlerFunc {
	return func(rc *web.RequestContext) (err error) {
		db, err := persistence.OpenMySQL()
		if err != nil {
			rc.Logger().Error("failed to connect mysql: %v", err)
			return
		}
		defer db.Close()

		tx, err := db.BeginTx(rc.Context(), nil)
		if err != nil {
			rc.Logger().Error("failed to begin transaction: %v", err)
			return
		}
		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
		persistence.SetTransaction(rc.Context(), tx)
		return next(rc)
	}
}
