package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql" // mysql
)

var (
	transactionPool map[context.Context]*sql.Tx = make(map[context.Context]*sql.Tx)
	mutex           sync.Mutex
)

// OpenMySQL ...
func OpenMySQL() (*sql.DB, error) {
	ds := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	return sql.Open("mysql", ds)
}

// SetTransaction ...
func SetTransaction(ctx context.Context, tx *sql.Tx) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := transactionPool[ctx]; ok {
		return
	}
	transactionPool[ctx] = tx
}

// TransactionFromContext ...
func TransactionFromContext(ctx context.Context) (*sql.Tx, error) {
	tx, ok := transactionPool[ctx]
	if !ok {
		return nil, fmt.Errorf("transaction is not found")
	}
	return tx, nil
}
