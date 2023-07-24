package database

import (
	"context"
)

type MockRow struct {
	MockScan func(dest ...any) error
}

func (mock *MockRow) Scan(dest ...any) error {
	return mock.MockScan(dest...)
}

func NewMockRow() *MockRow {
	return &MockRow{MockScan: func(dest ...any) error {
		return nil
	}}
}

type MockRows struct {
	MockClose func()
	MockErr   func() error
	MockNext  func() bool
	MockScan  func(dest ...any) error
}

func (mock *MockRows) Close() {
	mock.MockClose()
}

func (mock *MockRows) Err() error {
	return mock.MockErr()
}

func (mock *MockRows) Next() bool {
	return mock.MockNext()
}

func (mock *MockRows) Scan(dest ...any) error {
	return mock.MockScan(dest...)
}

func NewMockRows() *MockRows {
	return &MockRows{
		MockClose: func() {},
		MockErr:   func() error { return nil },
		MockNext:  func() bool { return true },
		MockScan:  func(dest ...any) error { return nil },
	}
}

type MockTransaction struct {
	MockRollback   func(ctx context.Context) error
	MockBulkInsert func(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error)
	MockCommit     func(ctx context.Context) error
}

func (mock *MockTransaction) Rollback(ctx context.Context) error {
	return mock.MockRollback(ctx)
}

func (mock *MockTransaction) BulkInsert(ctx context.Context, tableName string, columns []string,
	rows [][]any) (int, error) {
	return mock.MockBulkInsert(ctx, tableName, columns, rows)
}

func (mock *MockTransaction) Commit(ctx context.Context) error {
	return mock.MockCommit(ctx)
}

func NewMockTransaction() *MockTransaction {
	return &MockTransaction{
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

type MockDB struct {
	MockRow         *MockRow
	MockTransaction *MockTransaction
	MockRows        *MockRows
	MockClose       func()
	MockQueryRow    func(ctx context.Context, sql string, args ...any) Row
	MockBegin       func(ctx context.Context) (Transaction, error)
	MockQuery       func(ctx context.Context, sql string, args ...any) (Rows, error)
}

func (mock *MockDB) Close() {
	mock.MockClose()
}

func (mock *MockDB) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return mock.MockQueryRow(ctx, sql, args...)
}

func (mock *MockDB) Begin(ctx context.Context) (Transaction, error) {
	return mock.MockBegin(ctx)
}

func (mock *MockDB) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return mock.MockQuery(ctx, sql, args...)
}

func NewMockDB() *MockDB {
	mockRow := NewMockRow()
	mockTransaction := NewMockTransaction()
	mockRows := NewMockRows()
	return &MockDB{
		MockRow:         mockRow,
		MockTransaction: mockTransaction,
		MockRows:        mockRows,
		MockClose:       func() {},
		MockQueryRow: func(ctx context.Context, sql string, args ...any) Row {
			return mockRow
		},
		MockBegin: func(ctx context.Context) (Transaction, error) {
			return mockTransaction, nil
		},
		MockQuery: func(ctx context.Context, sql string, args ...any) (Rows, error) {
			return mockRows, nil
		},
	}
}
