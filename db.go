package server

import (
	"context"
	"database/sql"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/authservice/encryption"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

type Db interface {
	createUser()
}

// setDB sets initial config database
func setDb() (*sql.DB, error) {
	c, err := pgx.ParseConfig(os.Getenv("psqlEndpoint"))
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(c))
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	return db, nil
}

func (s *server) CreateUser(ctx context.Context, user User) (*sql.Rows, error) {
	encyptedPassword, err := encryption.Encrypt(user.Password)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	rows, err := sq.Insert("users").
		Columns("user_id", "name", "email", "age", "password").
		Suffix("RETURNING *").
		Values(user.Id, user.Name, user.Email, user.Age, encyptedPassword).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar).
		QueryContext(ctx)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return rows, nil
}
