package server

import (
	"net/http"
	"os"
	"path/filepath"
)

// Routes define the server routes
func (s *server) Routes() error {
	dir, err := os.Getwd()
	s.mux.HandleFunc("/health", s.handleHealth())
	s.mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(filepath.Join(dir, "/public")))))
	s.mux.HandleFunc("/login", s.handleLogin())
	s.mux.HandleFunc("/authenticate", s.handleAuthenticate())
	s.mux.HandleFunc("/signup", s.handleSignUp())
	s.mux.HandleFunc("/create-user", s.handleCreateUser())
	s.mux.HandleFunc("/.well-known/openid-configuration", s.handleOpenId(s.port))

	return err
}
