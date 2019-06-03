package filtercreator

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gemsorg/eligibility/pkg/filter"

	service "github.com/gemsorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.EligibilityService) http.Handler {
	return kithttp.NewServer(
		makeCreateFilterEndpoint(s),
		decodeFiltersResponse,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeFiltersResponse(_ context.Context, r *http.Request) (interface{}, error) {
	var f filter.Filter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&f)
	if err != nil {
		return nil, fmt.Errorf("missing params")
	}
	return f, nil
}
