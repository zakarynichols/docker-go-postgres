package http

import (
	"context"
	"net/http"

	"github.com/zakarynichols/parent-teacher-portal/postgresql"
)

func (s *Server) RegisterSchoolRoutes(ctx context.Context) {
	s.Router.Handle("/schools/new", handleCreateSchool(ctx, s.SchoolService))
}

func handleCreateSchool(ctx context.Context, sc SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := sc.CreateSchool(postgresql.School{Name: "New school for testing", Location: "1234 cool st.", Type: "public"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
	}
}
