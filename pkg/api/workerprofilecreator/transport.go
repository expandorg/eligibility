package workerprofilecreator

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/expandorg/eligibility/pkg/workerprofile"

	"github.com/expandorg/eligibility/pkg/apierror"

	service "github.com/expandorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.EligibilityService) http.Handler {
	return kithttp.NewServer(
		makeCreateWorkerProfileEndpoint(s),
		decodeNewWorkerProfileRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeNewWorkerProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var p workerprofile.NewProfile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}
	return p, nil
}
