package datastore

import (
	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetAllFilters() (filter.Filters, error)
	CreateFilter(filter.Filter) (filter.Filter, error)
	GetWorkerProfile(workerID string) (workerprofile.Profile, error)
}

type EligibilityStore struct {
	DB *sqlx.DB
}

func NewEligibilityStore(db *sqlx.DB) *EligibilityStore {
	return &EligibilityStore{
		DB: db,
	}
}

func (s *EligibilityStore) GetAllFilters() (filter.Filters, error) {
	filters := filter.Filters{}
	s.DB.Select(&filters, "SELECT * FROM filters")
	return filters, nil
}

func (s *EligibilityStore) CreateFilter(f filter.Filter) (filter.Filter, error) {
	result, err := s.DB.Exec("INSERT INTO filters (type, value) VALUES (?, ?)", f.Type, f.Value)
	if err != nil {
		return filter.Filter{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return filter.Filter{}, err
	}
	f.ID = uint64(id)
	return f, nil
}

func (s *EligibilityStore) GetWorkerProfile(workerID string) (workerprofile.Profile, error) {
	fs := filter.Filters{}
	attr := filter.GroupedFilters{}
	p := workerprofile.Profile{Attributes: attr}

	err := s.DB.Get(&p, "SELECT * FROM worker_profiles WHERE worker_id=? LIMIT 1", workerID)

	if err != nil {
		return p, err
	}

	err = s.DB.Select(&fs, "SELECT f.type, f.`value`, f.`id` from filters AS f inner join filters_workers AS fj on f.`id` = fj.`filter_id` where fj.`worker_id` = ?", workerID)
	if err != nil {
		return p, err
	}
	p.Attributes = fs.GroupByType()

	return p, nil
}

// func (s *EligibilityStore) GetJobFilters(jobID int)                                       {}
// func (s *EligibilityStore) GetJobWhiteList(jobID int)                                     {}
// func (s *EligibilityStore) CreateJobFilters(jobID int, filters []Filter, profile Profile) {}
// func (s *EligibilityStore) GetEligibleWorkerCount(filters []Filter)                       {}
// func (s *EligibilityStore) GetWorkerProfile(workerID int)                                 {}
// func (s *EligibilityStore) CreateWorkerProfile(workerID int, profile []Profile)           {}
// func (s *EligibilityStore) GetEligibleJobsForWorker(workerID int)                         {}
// func (s *EligibilityStore) IsWorkerEligibleForJob(workerID int, jobID int)                {}
