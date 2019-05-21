package main

import (
	"fmt"
	"net/http"

	"github.com/gemsorg/eligibility/log"

	"github.com/gemsorg/eligibility/pkg/server"
)

func main() {
	s := server.New()
	logger := log.New()
	logger.Log("info", fmt.Sprintf("Starting service on port 3000"))
	http.Handle("/", s)
	http.ListenAndServe(":3000", nil)
}
