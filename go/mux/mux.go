package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Mux struct {
	Router *mux.Router
}

func New() *Mux {
	r := mux.NewRouter()
	return &Mux{Router: r}
}

func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
