package main

import (
	"log"
	"net/http"

	server "github.com/authservice"
	"github.com/sirupsen/logrus"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal("could not start server")
	}
}

func run() error {
	const PORT = "8080"

	s := server.New()
	logrus.Infof("listening on http://localhost:%s\n", PORT)
	return http.ListenAndServe(":"+PORT, s.Mux())
}
