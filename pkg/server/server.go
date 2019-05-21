package server

import (
	"net/http"

	"github.com/expandorg/eligibility/pkg/api/eligibilityfetcher"
	"github.com/expandorg/eligibility/pkg/api/workerprofilecreator"
	"github.com/expandorg/eligibility/pkg/authentication"

	"github.com/expandorg/eligibility/pkg/api/workerprofilefetcher"

	"github.com/expandorg/eligibility/pkg/api/filtercreator"

	"github.com/jmoiron/sqlx"

	"github.com/expandorg/eligibility/pkg/api/filtersfetcher"

	"github.com/expandorg/eligibility/pkg/api/healthchecker"
	"github.com/expandorg/eligibility/pkg/service"
	"github.com/gorilla/mux"
)

func New(
	db *sqlx.DB,
	s service.EligibilityService,
) http.Handler {
	r := mux.NewRouter()

	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtersfetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/filters", filtercreator.MakeHandler(s)).Methods("POST")
	r.Handle("/workers/{worker_id}/profiles", workerprofilefetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/workers/{worker_id}/eligibility", eligibilityfetcher.MakeHandler(s)).Methods("GET")
	r.Handle("/workers/{worker_id}/profiles", workerprofilecreator.MakeHandler(s)).Methods("POST")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
