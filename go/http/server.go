package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zakarynichols/parent-teacher-portal/postgresql"
)

// Server extends the stdlib http.Server with the app's required services.
type Server struct {
	*http.Server
	*mux.Router
	UserService
	SchoolService
}

type UserService interface {
	QueryUsers() ([]postgresql.User, error)
}

type SchoolService interface {
	CreateSchool(postgresql.School) error
}
