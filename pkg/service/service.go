package service

import (
	"github.com/gemsorg/eligibility/pkg/datastore"
	"github.com/gemsorg/eligibility/pkg/workerprofile"

	"github.com/gemsorg/eligibility/pkg/filter"
)

type EligibilityService interface {
	Healthy() bool
	GetFilters() (filter.Filters, error)
	CreateFilter(filter.Filter) (filter.Filter, error)
	GetWorkerProfile(workerID string) (workerprofile.Profile, error)
}

type service struct {
	store datastore.Storage
}

func New(s datastore.Storage) *service {
	return &service{
		store: s,
	}
}

func (s *service) Healthy() bool {
	return true
}

func (s *service) GetFilters() (filter.Filters, error) {
	return s.store.GetAllFilters()
}

func (s *service) CreateFilter(f filter.Filter) (filter.Filter, error) {
	return s.store.CreateFilter(f)
}

func (s *service) GetWorkerProfile(workerID string) (workerprofile.Profile, error) {
	return s.store.GetWorkerProfile(workerID)
}

// GetWorkerProfile(workerID int)
// CreateWorkerProfile(workerID int, profile []Profile)
// GetJobFilters(jobID int)
// GetJobWhiteList(jobID int)
// CreateJobFilters(jobID int, filters []Filter, profile Profile)
// GetEligibleWorkerCount(filters []Filter)
// GetEligibleJobsForWorker(workerID int)
// IsWorkerEligibleForJob(workerID int, jobID int)
