package utils

import (
	"database/sql"
	"github.com/pkg/errors"
)

func SqlIsNotFount(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
