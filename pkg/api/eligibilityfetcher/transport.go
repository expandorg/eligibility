package eligibilityfetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/expandorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s service.EligibilityService) http.Handler {
	return kithttp.NewServer(
		makeEligibilityFetcherEndpoint(s),
		decodeEligibilityFetcherRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeEligibilityFetcherRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var ok bool
	workerID, ok := vars["worker_id"]
	if !ok {
		return nil, fmt.Errorf("missing worker_id parameter")
	}
	return WorkerEligibilityRequest{WorkerID: workerID}, nil
}
