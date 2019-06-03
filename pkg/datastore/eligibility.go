package datastore

import (
	"github.com/gemsorg/eligibility/pkg/filter"
)

type Driver interface {
	Select(dest interface{}, query string, args ...interface{}) error
}

type Storage interface {
	GetAllFilters() (filter.Filters, error)
}

type EligibilityStore struct {
	DB Driver
}

func NewEligibilityStore(db Driver) *EligibilityStore {
	return &EligibilityStore{
		DB: db,
	}
}

func (s *EligibilityStore) GetAllFilters() (filter.Filters, error) {
	filters := filter.Filters{}
	s.DB.Select(&filters, "SELECT * FROM filters")
	return filters, nil
}

// func (s *EligibilityStore) GetJobFilters(jobID int)                                       {}
// func (s *EligibilityStore) GetJobWhiteList(jobID int)                                     {}
// func (s *EligibilityStore) CreateJobFilters(jobID int, filters []Filter, profile Profile) {}
// func (s *EligibilityStore) GetEligibleWorkerCount(filters []Filter)                       {}
// func (s *EligibilityStore) GetWorkerProfile(workerID int)                                 {}
// func (s *EligibilityStore) CreateWorkerProfile(workerID int, profile []Profile)           {}
// func (s *EligibilityStore) GetEligibleJobsForWorker(workerID int)                         {}
// func (s *EligibilityStore) IsWorkerEligibleForJob(workerID int, jobID int)                {}
