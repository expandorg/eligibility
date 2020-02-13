package workerprofilefetcher

import (
	"context"

	"github.com/expandorg/eligibility/pkg/apierror"
	"github.com/expandorg/eligibility/pkg/authentication"
	service "github.com/expandorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeWorkerProfileFetcherEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, _ := authentication.ParseAuthData(ctx)
		svc.SetAuthData(data)
		req := request.(WorkerProfileRequest)
		p, err := svc.GetWorkerProfile(req.WorkerID)
		if err != nil {
			return p, errorResponse(err)
		}
		return p, nil
	}
}

type WorkerProfileRequest struct {
	WorkerID string
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(401, err.Error(), err)
}
