package filtercreator

import (
	"context"
	"log"

	"github.com/gemsorg/eligibility/pkg/filter"

	service "github.com/gemsorg/eligibility/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateFilterEndpoint(svc service.EligibilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(filter.Filter)
		saved, err := svc.CreateFilter(req)
		if err != nil {
			log.Fatal("error", err.Error())
			return nil, err
		}
		return saved, nil
	}
}
