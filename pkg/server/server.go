package server

import (
	"net/http"

	"github.com/gemsorg/eligibility/log"
	"github.com/gemsorg/eligibility/pkg/healthcheck"
	"github.com/gemsorg/eligibility/pkg/service"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	r := mux.NewRouter()
	l := log.New()
	hs := service.New(l)
	r.Handle("/_health", healthcheck.MakeHandler(hs, l)).Methods("GET")
	return r
}
