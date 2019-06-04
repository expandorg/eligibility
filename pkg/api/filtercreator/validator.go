package filtercreator

import (
	"github.com/gemsorg/eligibility/pkg/apierror"
	"github.com/gemsorg/eligibility/pkg/filter"
)

func validateRequest(req filter.Filter) (bool, *apierror.APIError) {
	var missingParams []string
	if req.Type == "" {
		missingParams = append(missingParams, "type")
	}
	if req.Value == "" {
		missingParams = append(missingParams, "value")
	}
	if len(missingParams) > 0 {
		return false, errorResponse(&apierror.MissingParameters{Params: missingParams})
	}
	return true, nil
}
