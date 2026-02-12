package db

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsUniqueErr(err error) bool {
	var targetErr *pgconn.PgError
	return errors.As(err, &targetErr) && targetErr.Code == "23505"
}
