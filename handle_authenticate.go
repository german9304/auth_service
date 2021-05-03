package server

import (
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type oauthQueryMeta struct {
	responseType string
	redirectUri  string
	clientId     string
	responseMode string
}

func newAuthQueryMeta(queryValues url.Values) oauthQueryMeta {
	return oauthQueryMeta{
		responseType: queryValues.Get("response_type"),
		redirectUri:  queryValues.Get("redirect_uri"),
		clientId:     queryValues.Get("client_id"),
		responseMode: queryValues.Get("response_mode"),
	}
}

type User struct {
	Name     string
	Password string
}

// handles authentication
func handleAuthenticate() http.HandlerFunc {
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
