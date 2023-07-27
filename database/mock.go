package database

import (
	"context"
)

type mockRow struct {
	MockScan func(dest ...any) error
}

func (mock *mockRow) Scan(dest ...any) error {
	return mock.MockScan(dest...)
}

func NewMockRow() mockRow {
	return mockRow{MockScan: func(dest ...any) error {
		return nil
	}}
}

type mockRows struct {
	MockClose func()
	MockErr   func() error
	MockNext  func() bool
	MockScan  func(dest ...any) error
}

func (mock *mockRows) Close() {
	mock.MockClose()
}

func (mock *mockRows) Err() error {
	return mock.MockErr()
}

func (mock *mockRows) Next() bool {
	return mock.MockNext()
}

func (mock *mockRows) Scan(dest ...any) error {
	return mock.MockScan(dest...)
}

func NewMockRows() mockRows {
	return mockRows{
		MockClose: func() {},
		MockErr:   func() error { return nil },
		MockNext:  func() bool { return true },
		MockScan:  func(dest ...any) error { return nil },
	}
}

type mockTransaction struct {
	MockRollback   func(ctx context.Context) error
	MockBulkInsert func(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error)
	MockCommit     func(ctx context.Context) error
}

func (mock *mockTransaction) Rollback(ctx context.Context) error {
	return mock.MockRollback(ctx)
}

func (mock *mockTransaction) BulkInsert(ctx context.Context, tableName string, columns []string,
	rows [][]any) (int, error) {
	return mock.MockBulkInsert(ctx, tableName, columns, rows)
}

func (mock *mockTransaction) Commit(ctx context.Context) error {
	return mock.MockCommit(ctx)
}

func NewMockTransaction() mockTransaction {
	return mockTransaction{
		MockRollback: func(ctx context.Context) error {
			return nil
		},
		MockBulkInsert: func(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error) {
			return 0, nil
		},
		MockCommit: func(ctx context.Context) error {
			return nil
		},
	}
}

type mockDB struct {
	MockRow         mockRow
	MockTransaction mockTransaction
	MockRows        mockRows
	MockClose       func()
	MockQueryRow    func(ctx context.Context, sql string, args ...any) Row
	MockBegin       func(ctx context.Context) (Transaction, error)
	MockQuery       func(ctx context.Context, sql string, args ...any) (Rows, error)
}

func (mock *mockDB) Close() {
	mock.MockClose()
}

func (mock *mockDB) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return mock.MockQueryRow(ctx, sql, args...)
}

func (mock *mockDB) Begin(ctx context.Context) (Transaction, error) {
	return mock.MockBegin(ctx)
}

func (mock *mockDB) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return mock.MockQuery(ctx, sql, args...)
}

func NewMockDB() mockDB {
	mockRow := NewMockRow()
	mockTransaction := NewMockTransaction()
	mockRows := NewMockRows()
	return mockDB{
		MockRow:         mockRow,
		MockTransaction: mockTransaction,
		MockRows:        mockRows,
		MockClose:       func() {},
		MockQueryRow: func(ctx context.Context, sql string, args ...any) Row {
			return &mockRow
		},
		MockBegin: func(ctx context.Context) (Transaction, error) {
			return &mockTransaction, nil
		},
		MockQuery: func(ctx context.Context, sql string, args ...any) (Rows, error) {
			return &mockRows, nil
		},
	}
}
