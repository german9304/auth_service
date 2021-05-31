package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/german9304/encryption"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// MockDatabaseQuery to mock D
type MockDatabaseQuery struct{}

func (md *MockDatabaseQuery) CreateUser(ctx context.Context, user User) (int64, error) {
	return 0, nil
}

func (md *MockDatabaseQuery) UserByEmail(ctx context.Context, email string) (*User, error) {
	return &User{}, nil
}

type DatabaseQuery interface {
	CreateUser(ctx context.Context, user User) (int64, error)
	UserByEmail(ctx context.Context, email string) (*User, error)
}

type Database struct {
	sql *sql.DB
}

// UserByEmail queries user by email
func (d *Database) UserByEmail(ctx context.Context, email string) (*User, error) {
	row, err := sq.Select("name", "email", "age", "password").
		From("users").
		Where(sq.Eq{"email": email}).
		RunWith(d.sql).
		PlaceholderFormat(sq.Dollar).
		QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	var user User

	for row.Next() {
		// if no rows with username
		err = row.Scan(&user.Name, &user.Email, &user.Age, &user.Password)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user with email %s", email)
		}
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a user and adds a row to the database
func (d *Database) CreateUser(ctx context.Context, user User) (int64, error) {
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
