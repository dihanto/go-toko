package database

import "context"

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Transaction interface {
	Rollback(context.Context) error
	Commit(ctx context.Context) error
}

type DB interface {
	Close()
	Exec(ctx context.Context, query string, args ...any) (interface{}, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Begin(ctx context.Context) (Transaction, error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
}
