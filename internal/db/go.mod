module github.com/german/db

replace github.com/german/encryption => ../encryption

replace github.com/german/user => ../user

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/german/encryption v0.0.0-00010101000000-000000000000
	github.com/german/user v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1
)
