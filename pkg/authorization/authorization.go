package authorization

import (
	"github.com/gemsorg/eligibility/pkg/authentication"
)

type Authorizer interface {
	CanAccessWorkerProfile(workerID uint64) (bool, error)
	SetAuthData(data authentication.AuthData)
}

type authorizor struct {
	authData authentication.AuthData
}

func NewAuthorizer() Authorizer {
	return &authorizor{
		authentication.AuthData{},
	}
}

func (a *authorizor) CanAccessWorkerProfile(workerID uint64) (bool, error) {
	if a.authData.UserID != workerID {
		return false, UnauthorizedAccess{}
	}
	return true, nil
}

func (a *authorizor) SetAuthData(data authentication.AuthData) {
	a.authData = data
}
