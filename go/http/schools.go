package http

import (
	"context"
	"encoding/json"
	"net/http"

	ptp "github.com/zakarynichols/parent-teacher-portal"
	"github.com/zakarynichols/parent-teacher-portal/mux"
)

func (s *Server) RegisterSchoolRoutes(ctx context.Context) {
	s.mux.Handle("/schools", handleCreateSchool(ctx, s.schoolService), "POST")
	s.mux.Handle("/schools/{id}", handleGetSchool(ctx, s.schoolService), "GET")
	s.mux.Handle("/schools", handleGetAllSchools(ctx, s.schoolService), "GET")
	s.mux.Handle("/schools/{id}", handleUpdateSchool(ctx, s.schoolService), "PUT")
	s.mux.Handle("/schools/{id}", handleDeleteSchool(ctx, s.schoolService), "DELETE")
}

func handleCreateSchool(ctx context.Context, ss ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var school ptp.School

		if err := json.NewDecoder(r.Body).Decode(&school); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := ss.CreateSchool(school)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
	}
}

func handleGetSchool(ctx context.Context, ss ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "ID not found in URL", http.StatusBadRequest)
			return
		}

		school, err := ss.GetSchool(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(school)
	}
}

func handleGetAllSchools(ctx context.Context, ss ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		schools, err := ss.GetAllSchools()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(schools)
	}
}

func handleUpdateSchool(ctx context.Context, ss ptp.SchoolService) http.HandlerFunc {
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

		err := ss.UpdateSchool(id, school)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func handleDeleteSchool(ctx context.Context, ss ptp.SchoolService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		err := ss.DeleteSchool(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
