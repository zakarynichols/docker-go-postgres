package http

import (
	"context"
	"encoding/json"
	"net/http"

	ptp "github.com/zakarynichols/parent-teacher-portal"
	"github.com/zakarynichols/parent-teacher-portal/mux"
)

func (s *Server) RegisterSchoolRoutes(ctx context.Context) {
	s.mux.Router.Handle("/schools", handleCreateSchool(ctx, s.schoolService)).Methods("POST")
	s.mux.Router.Handle("/schools/{id}", handleGetSchool(ctx, s.schoolService)).Methods("GET")
	s.mux.Router.Handle("/schools", handleGetAllSchools(ctx, s.schoolService)).Methods("GET")
	s.mux.Router.Handle("/schools/{id}", handleUpdateSchool(ctx, s.schoolService)).Methods("PUT")
	s.mux.Router.Handle("/schools/{id}", handleDeleteSchool(ctx, s.schoolService)).Methods("DELETE")
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

func handleGetSchool(ctx context.Context, sc ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "ID not found in URL", http.StatusBadRequest)
			return
		}

		school, err := sc.GetSchool(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(school)
	}
}

func handleGetAllSchools(ctx context.Context, sc ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		schools, err := sc.GetAllSchools()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(schools)
	}
}

func handleUpdateSchool(ctx context.Context, sc ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "ID not found in URL", http.StatusBadRequest)
			return
		}

		var school ptp.School
		if err := json.NewDecoder(r.Body).Decode(&school); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := sc.UpdateSchool(id, school)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func handleDeleteSchool(ctx context.Context, sc ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		err := sc.DeleteSchool(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
