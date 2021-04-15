package main

import (
	"auth_service/routes"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Printf("route: %s\n", routes.Route())
}
