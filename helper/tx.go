package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx, err *error) {
	if *err != nil {
		errRolback := tx.Rollback()
		if errRolback != nil {
			log.Println(err)
		}
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			log.Println(err)
		}
	}
}
