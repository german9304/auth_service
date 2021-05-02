package server

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type authQueryMeta struct {
	responseType string
	redirectUri  string
	clientId     string
	responseMode string
}

func (r authQueryMeta) new(queryValues url.Values) authQueryMeta {
	return authQueryMeta{
		responseType: queryValues.Get("response_type"),
		redirectUri:  queryValues.Get("redirect_uri"),
		clientId:     queryValues.Get("client_id"),
		responseMode: queryValues.Get("response_mode"),
	}
}

// handles authentication
func handleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("user is authenticating")
		http.ServeFile(w, r, filepath.Join(dir, "/client", "/public", "/index.html"))
	}
}
