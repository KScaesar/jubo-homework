package database

import (
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

func PgsqlError(err *pgconn.PgError) error {
	switch {
	case err.Code == "23505":
		return errors.ErrKeyDuplicated
	default:
		return errors.Join3rdParty(errors.ErrSystem, err)
	}
}
