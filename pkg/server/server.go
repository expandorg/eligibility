package server

import (
	"net/http"

	"github.com/gemsorg/eligibility/pkg/authorization"

	"github.com/gemsorg/eligibility/pkg/api/workerprofilecreator"
	"github.com/gemsorg/eligibility/pkg/authentication"

	"github.com/gemsorg/eligibility/pkg/api/workerprofilefetcher"

	"github.com/gemsorg/eligibility/pkg/api/filtercreator"

	"github.com/jmoiron/sqlx"

	"github.com/gemsorg/eligibility/pkg/datastore"

	"github.com/gemsorg/eligibility/pkg/api/filtersfetcher"

	"github.com/gemsorg/eligibility/pkg/api/healthchecker"
	"github.com/gemsorg/eligibility/pkg/service"
	"github.com/gorilla/mux"
)

func New(db *sqlx.DB) http.Handler {
	r := mux.NewRouter()
	ds := datastore.NewEligibilityStore(db)
	authorizer := authorization.NewAuthorizer()
	s := service.New(ds, authorizer)
	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtersfetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtercreator.MakeHandler(s)).Methods("POST")
	r.Handle("/workers/{worker_id}/profiles", workerprofilefetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/workers/{worker_id}/profiles", workerprofilecreator.MakeHandler(s)).Methods("POST")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
