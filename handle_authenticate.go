package server

import (
	"context"
	"net/http"
	"net/url"

	"github.com/german9304/encryption"
	"github.com/sirupsen/logrus"
)

// handles authentication authenticates user
func (s *server) handleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()
		reqCtx := r.WithContext(ctx)
		logrus.Info("authenticating with username and password")
		// parse url encoded form
		err := reqCtx.ParseForm()
		if err != nil {
			logrus.Info(err)
		}
		email := reqCtx.Form.Get("email")
		password := reqCtx.Form.Get("password")

		user, err := s.db.UserByEmail(ctx, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		decryptedPassword := encryption.Decrypt(user.Password)
		if decryptedPassword != password {
			http.Error(w, "invalid password", http.StatusBadRequest)
			return
		}

		redirectUri := r.URL.Query().Get("redirectUri")
		authorizationEndpoint := r.URL.Query().Get("authorizationEndpoint")
		responseMode := r.URL.Query().Get("responseMode")

		logrus.Infof("authorization endpoint: %s\n", authorizationEndpoint)
		if responseMode == "form_post" {
			values := url.Values{}
			values.Add("id_token", "")
			w.Header().Set("Content-type", "application/x-www-form-urlencoded")
			w.Write([]byte(values.Encode()))
		}

		http.Redirect(w, r, redirectUri, http.StatusFound)
	}
}
