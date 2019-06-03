package filtersfetcher

import (
	"context"

	"github.com/gemsorg/eligibility/pkg/filter"

	service "github.com/gemsorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeFiltersFetcherEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		filters, _ := svc.GetFilters()
		grouped := filters.GroupByType()
		return FiltersResponse{grouped}, nil
	}
}

type FiltersResponse struct {
	Filters filter.GroupedFilters `json:"filters"`
}
