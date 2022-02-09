package database

import "database/sql"

// Verify if has no rows affected or if has some error
// if has some error return the error, if has no rows return ErrNoRows
// else return nil
func CheckHasNoRowsAffected(res sql.Result) error {
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows <= 0 {
		return sql.ErrNoRows
	}

	return nil
}
