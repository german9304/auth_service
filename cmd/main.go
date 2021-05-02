package main

import (
	"log"
	"net/http"

	server "github.com/authservice"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = run()
	if err != nil {
		log.Fatal("could not start server")
	}
}

func run() error {
	const PORT = "8081"

	s := server.New(PORT)
	logrus.Infof("listening on http://localhost:%s\n", PORT)
	return http.ListenAndServe(":"+PORT, s.Mux())
}
