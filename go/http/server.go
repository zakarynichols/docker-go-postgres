package http

import (
	"net/http"

	"github.com/gorilla/mux"
	ptp "github.com/zakarynichols/parent-teacher-portal"
)

// Server extends the stdlib http.Server with the app's required services.
type Server struct {
	*http.Server
	*mux.Router
	ptp.SchoolService
}
