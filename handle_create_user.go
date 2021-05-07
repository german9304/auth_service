package server

import "net/http"

func handleCreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-type", "application-json")
	}
}
