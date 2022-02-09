package database

import (
	"database/sql"
	"errors"
)

// IsNoRowError identifies a row error type
// It allow both directly identification and unwrap one
// for those who wraps error using fmt.Errorf("... %w ...")
func IsNoRowsError(err error) bool {
	if ok := errors.Is(err, sql.ErrNoRows); ok {
		return ok
	}

	return errors.Is(errors.Unwrap(err), sql.ErrNoRows)
}
