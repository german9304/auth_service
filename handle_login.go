package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func (s *server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			logrus.Fatal(err)
		}
		http.ServeFile(w, r, filepath.Join(dir, "/public", "/login.html"))
	}
}
