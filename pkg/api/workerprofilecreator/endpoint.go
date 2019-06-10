package workerprofilecreator

import (
	"context"

	"github.com/gemsorg/eligibility/pkg/apierror"
	"github.com/gemsorg/eligibility/pkg/workerprofile"

	service "github.com/gemsorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateWorkerProfileEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(workerprofile.NewProfile)
		saved, err := svc.CreateWorkerProfile(req)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

type WorkerProfileResponse struct {
	Profile workerprofile.Profile `json:"profile"`
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
