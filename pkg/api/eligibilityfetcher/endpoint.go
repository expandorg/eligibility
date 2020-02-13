package eligibilityfetcher

import (
	"context"

	"github.com/expandorg/eligibility/pkg/apierror"
	service "github.com/expandorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeEligibilityFetcherEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WorkerEligibilityRequest)
		el, err := svc.GetWorkerEligibility(req.WorkerID)
		if err != nil {
			return nil, errorResponse(err)
		}
		return el, nil
	}
}

type WorkerEligibilityRequest struct {
	WorkerID string
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
