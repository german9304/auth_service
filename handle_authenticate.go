package server

import (
	"net/http"
	"net/url"
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

	}
}
