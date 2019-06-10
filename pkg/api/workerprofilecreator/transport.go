package workerprofilecreator

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gemsorg/eligibility/pkg/workerprofile"

	"github.com/gemsorg/eligibility/pkg/apierror"

	service "github.com/gemsorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.EligibilityService) http.Handler {
	return kithttp.NewServer(
		makeCreateWorkerProfileEndpoint(s),
		decodeFiltersRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeFiltersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var p workerprofile.NewProfile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}
	if valid, err := validateRequest(p); !valid {
		return nil, err
	}

	return p, nil
}
