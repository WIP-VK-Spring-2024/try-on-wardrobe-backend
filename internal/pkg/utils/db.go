package utils

import (
	"errors"

	"try-on/internal/pkg/app_errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func GormError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return errors.Join(err, app_errors.ErrNotFound)
	case gorm.ErrDuplicatedKey:
		return errors.Join(err, app_errors.ErrAlreadyExists)
	case gorm.ErrInvalidField:
		return errors.Join(err, app_errors.ErrBadRequest)
	default:
		return err
	}
}

func PgxError(err error) error {
	if err == pgx.ErrNoRows {
		return app_errors.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return err
	}

	var appError error

	switch pgErr.Code {
	case pgerrcode.UniqueViolation:
		appError = app_errors.ErrAlreadyExists
	case pgerrcode.NoData:
		appError = app_errors.ErrNotFound
	case pgerrcode.ForeignKeyViolation:
		appError = app_errors.ErrNoRelatedEntity
	case pgerrcode.IntegrityConstraintViolation:
		appError = app_errors.ErrConstraintViolation
	}

	return errors.Join(err, appError)
}

func TranslateGormError[T any](item *T, err error) (*T, error) {
	err = GormError(err)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func TranslatePgxError[T any](item *T, err error) (*T, error) {
	err = PgxError(err)
	if err != nil {
		return nil, err
	}
	return item, nil
}
