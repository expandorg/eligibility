package service

import (
	"strconv"

	"github.com/gemsorg/eligibility/pkg/authentication"
	"github.com/gemsorg/eligibility/pkg/authorization"
	"github.com/gemsorg/eligibility/pkg/datastore"
	"github.com/gemsorg/eligibility/pkg/workerprofile"

	"github.com/gemsorg/eligibility/pkg/filter"
)

type EligibilityService interface {
	Healthy() bool
	GetFilters() (filter.Filters, error)
	CreateFilter(filter.Filter) (filter.Filter, error)
	GetWorkerProfile(workerID string) (workerprofile.Profile, error)
	CreateWorkerProfile(workerprofile.NewProfile) (workerprofile.Profile, error)
	SetAuthData(data authentication.AuthData)
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

// GetWorkerProfile(workerID int)
// CreateWorkerProfile(workerID int, profile []Profile)
// GetJobFilters(jobID int)
// GetJobWhiteList(jobID int)
// CreateJobFilters(jobID int, filters []Filter, profile Profile)
// GetEligibleWorkerCount(filters []Filter)
// GetEligibleJobsForWorker(workerID int)
// IsWorkerEligibleForJob(workerID int, jobID int)
