package db

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/german/encryption"
	model "github.com/german/user"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DatabaseQuery interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
}

type Database struct {
	sql *sql.DB
}

// CreateUser creates a user and adds a row to the database
func (d *Database) CreateUser(ctx context.Context, user model.User) (int64, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	encyptedPassword, err := encryption.Encrypt(user.Password)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	rows, err := sq.Insert("users").
		Columns("user_id", "name", "email", "age", "password").
		Suffix("RETURNING *").
		Values(fmt.Sprintf("PROD-%s", uuid), user.Name, user.Email, user.Age, encyptedPassword).
		RunWith(d.sql).
		PlaceholderFormat(sq.Dollar).
		ExecContext(ctx)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	row, err := rows.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	return row, nil
}
