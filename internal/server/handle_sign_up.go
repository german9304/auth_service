package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func (s *server) handleSignUp() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		logrus.Info("handle signup")
		dir, err := os.Getwd()
		if err != nil {
			logrus.Fatal(err)
		}
		http.ServeFile(rw, r, filepath.Join(dir, "/public", "/signup.html"))
	}
}
