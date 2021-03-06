package service

import (
	"strconv"

	"github.com/expandorg/eligibility/pkg/eligibility"

	"github.com/expandorg/eligibility/pkg/authentication"
	"github.com/expandorg/eligibility/pkg/authorization"
	"github.com/expandorg/eligibility/pkg/datastore"
	"github.com/expandorg/eligibility/pkg/workerprofile"

	"github.com/expandorg/eligibility/pkg/filter"
)

type EligibilityService interface {
	Healthy() bool
	GetFilters() (filter.Filters, error)
	CreateFilter(filter.Filter) (filter.Filter, error)
	GetWorkerProfile(workerID string) (workerprofile.Profile, error)
	CreateWorkerProfile(workerprofile.NewProfile) (workerprofile.Profile, error)
	SetAuthData(data authentication.AuthData)
	GetWorkerEligibility(workerID string) (eligibility.WorkerEligibility, error)
}

type service struct {
	store      datastore.Storage
	authorizor authorization.Authorizer
}

func New(s datastore.Storage, a authorization.Authorizer) *service {
	return &service{
		store:      s,
		authorizor: a,
	}
}

func (s *service) Healthy() bool {
	return true
}

func (s *service) GetFilters() (filter.Filters, error) {
	return s.store.GetAllFilters(([]int{}))
}

func (s *service) CreateFilter(f filter.Filter) (filter.Filter, error) {
	return s.store.CreateFilter(f)
}

func (s *service) GetWorkerProfile(workerID string) (workerprofile.Profile, error) {
	authUserID, _ := strconv.ParseUint(workerID, 10, 64)

	_, err := s.authorizor.CanAccessWorkerProfile(authUserID)
	if err != nil {
		return workerprofile.Profile{}, err
	}
	return s.store.GetWorkerProfile(workerID)
}

func (s *service) CreateWorkerProfile(wp workerprofile.NewProfile) (workerprofile.Profile, error) {
	_, err := s.authorizor.CanAccessWorkerProfile(wp.WorkerID)
	if err != nil {
		return workerprofile.Profile{}, err
	}
	return s.store.CreateWorkerProfile(wp)
}

func (s *service) SetAuthData(data authentication.AuthData) {
	s.authorizor.SetAuthData(data)
}

func (s *service) GetWorkerEligibility(workerID string) (eligibility.WorkerEligibility, error) {
	var profileComplete bool
	w, j, err := s.store.GetWorkerEligibility(workerID)
	if err != nil {
		return eligibility.WorkerEligibility{}, err
	}

	// Worker doesn't have country set in their profile
	if w.FilterID != 0 {
		profileComplete = true
	}
	return eligibility.GetWorkerEligibility(w, j, profileComplete), nil
}
