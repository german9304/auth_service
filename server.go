package server

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

type server struct {
	mux  *http.ServeMux
	port string
	db   DatabaseQuery
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) Mux() *http.ServeMux {
	return s.mux
}

// newDatabase sets initial config database
func newDatabase() (DatabaseQuery, error) {
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
	return &Database{sql: db}, nil
}

func New(port string) (*server, error) {
	db, err := newDatabase()
	if err != nil {
		return nil, err
	}
	s := &server{
		mux:  http.NewServeMux(),
		port: port,
		db:   db,
	}
	s.Routes()
	return s, nil
}
