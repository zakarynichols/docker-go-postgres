package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Mux struct {
	router *mux.Router
}

func New() *Mux {
	r := mux.NewRouter()
	return &Mux{router: r}
}

func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func (m Mux) Handle(path string, handler http.Handler, method string) {
	m.router.Handle(path, handler).Methods(method)
}
