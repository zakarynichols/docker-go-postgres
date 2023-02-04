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

type Config struct {
	Addr          string
	SchoolService ptp.SchoolService
}

func New(config Config) *Server {
	m := mux.New()
	c := cors.New()
	h := c.Handler(m)
	return &Server{
		server: &http.Server{
			Addr:    ":" + config.Addr,
			Handler: h,
		},
		mux:           m,
		schoolService: config.SchoolService,
	}
}

func (s Server) ListenTLS(cert, key string) error {
	return s.server.ListenAndServeTLS(cert, key)
}

func (s Server) Listen() error {
	return s.server.ListenAndServe()
}
