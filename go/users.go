package app

import "go-postgres-docker/postgresql"

type UserService interface {
	QueryUsers() ([]postgresql.User, error)
}
