package filtercreator

import (
	"context"

	"github.com/gemsorg/eligibility/pkg/apierror"
	"github.com/gemsorg/eligibility/pkg/filter"

	service "github.com/gemsorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateFilterEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(filter.Filter)
		saved, err := svc.CreateFilter(req)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
