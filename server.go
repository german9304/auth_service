package server

import (
	"net/http"
)

type server struct {
	mux *http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) Mux() *http.ServeMux {
	return s.mux
}

func New() *server {
	s := &server{
		mux: http.NewServeMux(),
	}
	s.Routes()
	return s
}
