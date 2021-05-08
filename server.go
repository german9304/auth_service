package server

import (
	"database/sql"
	"net/http"
)

type server struct {
	mux  *http.ServeMux
	port string
	db   *sql.DB
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) Mux() *http.ServeMux {
	return s.mux
}

func New(port string) (*server, error) {
	db, err := setDb()
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
