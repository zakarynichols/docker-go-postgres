package http

import (
	"net/http"

	ptp "github.com/zakarynichols/parent-teacher-portal"
	"github.com/zakarynichols/parent-teacher-portal/cors"
	"github.com/zakarynichols/parent-teacher-portal/mux"
)

type Server struct {
	server        *http.Server
	mux           *mux.Mux
	schoolService ptp.SchoolService
}

func New(addr string, ss ptp.SchoolService) *Server {
	m := mux.New()
	c := cors.New()

	return &Server{
		server: &http.Server{
			Addr:    ":" + addr,
			Handler: c.Handler(m),
		},
		mux:           m,
		schoolService: ss,
	}
}

func (s Server) ListenTLS() error {
	return s.server.ListenAndServeTLS("cert.pem", "key.pem")
}

func (s Server) Listen() error {
	return s.server.ListenAndServe()
}
