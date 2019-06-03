package server

import (
	"net/http"

	"github.com/gemsorg/eligibility/pkg/datastore"

	"github.com/gemsorg/eligibility/pkg/api/filtersfetcher"

	"github.com/gemsorg/eligibility/pkg/api/healthchecker"
	"github.com/gemsorg/eligibility/pkg/service"
	"github.com/gorilla/mux"
)

func New(db datastore.Driver) http.Handler {
	r := mux.NewRouter()
	ds := datastore.NewEligibilityStore(db)
	s := service.New(ds)
	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtersfetcher.MakeHandler(s)).Methods("GET")
	return r
}
