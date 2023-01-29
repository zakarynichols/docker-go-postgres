package http

import (
	"context"
	"encoding/json"
	"net/http"

	ptp "github.com/zakarynichols/parent-teacher-portal"
)

func (s *Server) RegisterSchoolRoutes(ctx context.Context) {
	s.Router.Handle("/schools", handleCreateSchool(ctx, s.SchoolService)).Methods("POST")
}

func handleCreateSchool(ctx context.Context, sc ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var school ptp.School

		if err := json.NewDecoder(r.Body).Decode(&school); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := sc.CreateSchool(school)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
	}
}
