package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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

type User struct {
	Name     string
	Password string
}

// handles authentication
func handleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "could not read request body", http.StatusBadRequest)
			return
		}

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Fatal(err)
		}
		logrus.Infof("name: %s, password: %s\n", user.Name, user.Password)

		w.WriteHeader(http.StatusOK)
	}
}
