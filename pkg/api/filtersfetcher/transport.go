package filtersfetcher

import (
	"context"
	"encoding/json"
	"net/http"

	service "github.com/gemsorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.EligibilityService) http.Handler {
	return kithttp.NewServer(
		makeFiltersFetcherEndpoint(s),
		decodeFiltersResponse,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeFiltersResponse(_ context.Context, r *http.Request) (interface{}, error) {
	return FiltersResponse{}, nil
}
