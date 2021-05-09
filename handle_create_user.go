package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// handleCreateUser handler to create a user
func handleCreateUser(s DatabaseQuery) http.HandlerFunc {
	type Response struct {
		Data    string `json:"data"`
		Created bool   `json:"created"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-type", "application-json")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Error(err)
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		var user User
		err = json.Unmarshal(b, &user)
		if err != nil {
			logrus.Error(err)
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.TODO()
		_, err = s.CreateUser(ctx, user)
		if err != nil {
			logrus.Info("error here create user")
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		logrus.Info("user is created")

		response := Response{Data: "user created", Created: true}
		responseBody, err := json.Marshal(response)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write(responseBody)
	}
}
