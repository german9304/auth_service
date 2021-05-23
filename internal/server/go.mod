module github.com/german/server

go 1.16

require (
	github.com/jackc/pgx/v4 v4.11.0
	github.com/sirupsen/logrus v1.8.1
)

replace github.com/german/db => ../db

replace github.com/german/encryption => ../encryption
