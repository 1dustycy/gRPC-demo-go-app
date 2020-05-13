package web

import (
	"net/http"

	"github.com/betterchen/go-project-tmpl/pkg/multiservices"
	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	r.Path("/healthz").Handler(healthz())
}

func healthz() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if multiservices.Healthz() {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("ok"))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
		},
	)
}
