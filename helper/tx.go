package helper

import (
	"database/sql"

	"github.com/dihanto/go-toko/exception"
)

func CommitOrRollback(tx *sql.Tx, err *error) {
	if *err != nil {
		errRolback := tx.Rollback()
		if errRolback != nil {
			exception.ErrorHandler(nil, nil, errRolback)
		}
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			exception.ErrorHandler(nil, nil, errCommit)
		}
	}
}
