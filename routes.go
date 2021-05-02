package server

import (
	"net/http"
	"os"
	"path/filepath"
)

func (s *server) Routes() error {
	dir, err := os.Getwd()
	s.mux.HandleFunc("/health", handleHealth())
	s.mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(filepath.Join(dir, "/client", "/public")))))
	s.mux.HandleFunc("/login", handleLogin())
	s.mux.HandleFunc("/authenticate", handleAuthenticate())
	s.mux.HandleFunc("/.well-known/openid-configuration", handleOpenId())

	return err
}
