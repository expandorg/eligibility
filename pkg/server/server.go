package server

import (
	"net/http"
	"os"

	"github.com/gemsorg/eligibility/pkg/api/workerprofilecreator"

	"github.com/gemsorg/eligibility/pkg/api/workerprofilefetcher"

	"github.com/gemsorg/eligibility/pkg/api/filtercreator"

	"github.com/jmoiron/sqlx"

	"github.com/gemsorg/eligibility/pkg/datastore"

	"github.com/gemsorg/eligibility/pkg/api/filtersfetcher"

	"github.com/gemsorg/eligibility/pkg/api/healthchecker"
	"github.com/gemsorg/eligibility/pkg/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New(db *sqlx.DB) http.Handler {
	r := mux.NewRouter()
	ds := datastore.NewEligibilityStore(db)
	s := service.New(ds)
	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtersfetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtercreator.MakeHandler(s)).Methods("POST")
	r.Handle("/profiles/{worker_id}", workerprofilefetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/profiles/{worker_id}", workerprofilecreator.MakeHandler(s)).Methods("POST")
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
	return loggedRouter
}
