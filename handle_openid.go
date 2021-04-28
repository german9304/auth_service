package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type openIdFields struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	Issuer                string `json:"issuer"`
	JwksUri               string `json:"jwks_uri"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}

// handleOpenId endpoint for openId meta fields
// to authenticate users.
func handleOpenId() http.HandlerFunc {
	const uri = "http://localhost:"
	return func(w http.ResponseWriter, r *http.Request) {
		openIdField := openIdFields{
			AuthorizationEndpoint: fmt.Sprintf("%s%d/authorize", uri, PORT),
			Issuer:                fmt.Sprintf("%s%d", uri, PORT),
			JwksUri:               fmt.Sprintf("%s%d/.well-known/jwks.json", uri, PORT),
			TokenEndpoint:         fmt.Sprintf("%s%d/token", uri, PORT),
			UserinfoEndpoint:      fmt.Sprintf("%s%d/userinfo", uri, PORT),
		}

		b, err := json.Marshal(openIdField)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logrus.Info("fetching openid configuration")

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
