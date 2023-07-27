package database

import (
	"context"
	"log"

	"github.com/dihanto/go-toko/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transaction struct {
	tx pgx.Tx
}

func NewTransaction(tx pgx.Tx) Transaction {
	return &transaction{tx: tx}
}

func (postgres *transaction) Rollback(ctx context.Context) error {
	return postgres.tx.Rollback(ctx)
}

func (postgres *transaction) Commit(ctx context.Context) error {
	return postgres.tx.Commit(ctx)
}

type postgresDB struct {
	pool *pgxpool.Pool
}

func (postgres *postgresDB) Close() {
	postgres.pool.Close()
}

func (postgres *postgresDB) Begin(ctx context.Context) (Transaction, error) {
	pgTx, err := postgres.pool.Begin(ctx)
	tx := NewTransaction(pgTx)
	return tx, err
}

func (postgres *postgresDB) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return postgres.pool.QueryRow(ctx, sql, args...)
}

func (postgres *postgresDB) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return postgres.pool.Query(ctx, sql, args...)
}

func (postgres *postgresDB) Exec(ctx context.Context, sql string, args ...any) (any, error) {
	return postgres.pool.Exec(ctx, sql, args...)
}

func NewDB(conf config.Config) (DB, error) {
	pool, err := pgxpool.New(context.Background(), conf.DB)
	if err != nil {
		log.Println(err)
	}
	db := &postgresDB{pool: pool}

	return db, nil
}
