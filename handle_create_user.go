package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func handleCreateUser(s *server) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-type", "application-json")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Fatal(err)
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
		result, err := s.CreateUser(ctx, user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		result.Scan(&user)

		logrus.Infof("%v\n", user)

		b, err = json.Marshal(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		rw.Write(b)
	}
}
