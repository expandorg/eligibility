package filtersfetcher

import (
	"context"

	"github.com/expandorg/eligibility/pkg/apierror"
	service "github.com/expandorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeFiltersFetcherEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		filters, err := svc.GetFilters()
		if err != nil {
			return nil, errorResponse(err)
		}

		grouped := filters.GroupByType()
		return grouped, nil
	}
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
