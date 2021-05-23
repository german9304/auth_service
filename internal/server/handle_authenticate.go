package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// handles authentication
func (s *server) handleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("authenticating with username and password")
		// parse url encoded form
		err := r.ParseForm()
		if err != nil {
			logrus.Info(err)
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		logrus.Infof("password: %s\n", password)
		logrus.Infof("username: %s\n", username)

		http.Redirect(w, r, "http://localhost:8080/api/health", http.StatusFound)
	}
}
