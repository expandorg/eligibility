package filtercreator

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gemsorg/eligibility/pkg/apierror"

	"github.com/gemsorg/eligibility/pkg/filter"

	service "github.com/gemsorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.EligibilityService) http.Handler {
	return kithttp.NewServer(
		makeCreateFilterEndpoint(s),
		decodeFiltersRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeFiltersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var f filter.Filter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&f)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}
	if valid, err := validateRequest(f); !valid {
		return nil, err
	}
	return f, nil
}
