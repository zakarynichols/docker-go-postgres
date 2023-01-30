package http

import (
	"net/http"

	"github.com/gorilla/mux"
	ptp "github.com/zakarynichols/parent-teacher-portal"
	"github.com/zakarynichols/parent-teacher-portal/cors"
)

// Server extends the stdlib http.Server with the app's required services.
type Server struct {
	server        *http.Server
	router        *mux.Router
	schoolService ptp.SchoolService
}

func New(addr string, ss ptp.SchoolService) *Server {
	r := mux.NewRouter()
	c := cors.New()

	return &Server{
		server: &http.Server{
			Addr:    ":" + addr,
			Handler: c.Handler(r),
		},
		router:        r,
		schoolService: ss,
	}
}

func (s *Server) ListenTLS() error {
	return s.server.ListenAndServeTLS("cert.pem", "key.pem")
}
