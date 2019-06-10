package workerprofilecreator

import (
	"github.com/gemsorg/eligibility/pkg/apierror"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
)

func validateRequest(req workerprofile.NewProfile) (bool, *apierror.APIError) {
	return true, nil
}
