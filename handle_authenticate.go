package server

import "net/http"

// handles authentication
func handleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
