package server

import (
	"net/http"
)

type server struct {
	mux  *http.ServeMux
	port string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) Mux() *http.ServeMux {
	return s.mux
}

func New(port string) *server {
	s := &server{
		mux:  http.NewServeMux(),
		port: port,
	}
	s.Routes()
	return s
}
