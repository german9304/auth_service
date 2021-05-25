module github.com/authservice

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/german9304/encryption v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.2.0
	github.com/jackc/pgx/v4 v4.11.0
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.8.1

)

replace github.com/german9304/encryption => ./internal/encryption
