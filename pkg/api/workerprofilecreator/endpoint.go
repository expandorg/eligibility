package workerprofilecreator

import (
	"context"

	"github.com/expandorg/eligibility/pkg/apierror"
	"github.com/expandorg/eligibility/pkg/authentication"
	"github.com/expandorg/eligibility/pkg/workerprofile"

	service "github.com/expandorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateWorkerProfileEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, _ := authentication.ParseAuthData(ctx)
		svc.SetAuthData(data)
		req := request.(workerprofile.NewProfile)
		saved, err := svc.CreateWorkerProfile(req)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
