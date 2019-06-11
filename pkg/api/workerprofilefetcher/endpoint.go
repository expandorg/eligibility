package workerprofilefetcher

import (
	"context"

	service "github.com/gemsorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeWorkerProfileFetcherEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(WorkerProfileRequest)
		p, _ := svc.GetWorkerProfile(req.WorkerID)
		return p, nil
	}
}

type WorkerProfileRequest struct {
	WorkerID string
}
