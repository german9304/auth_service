package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type health struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		greet := health{Data: "Health endpoint", Status: "Ok"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		b, err := json.Marshal(greet)
		if err != nil {
			logrus.Error("error marshaling struct")
		}
		logrus.Info("health endpoint")
		w.Write(b)
	}
}
