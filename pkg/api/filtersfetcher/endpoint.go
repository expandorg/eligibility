package filtersfetcher

import (
	"context"

	service "github.com/gemsorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeFiltersFetcherEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		filters, _ := svc.GetFilters()
		grouped := filters.GroupByType()
		return grouped, nil
	}
}
