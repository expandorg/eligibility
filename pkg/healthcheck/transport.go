package healthcheck

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gemsorg/eligibility/log"
	service "github.com/gemsorg/eligibility/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s service.EligibilityService, log log.Logger) http.Handler {
	return kithttp.NewServer(
		makeHealthyEndpoint(s),
		decodeHealthRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return HealthyResponse{}, nil
}
