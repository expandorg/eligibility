package datastore

import (
	"log"

	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/jmoiron/sqlx"
)

type EligibilityStore struct {
	db *sqlx.DB
}

func NewEligibilityStore(db *sqlx.DB) *EligibilityStore {
	return &EligibilityStore{
		db: db,
	}
}

func (s *EligibilityStore) GetAllFilters() (filter.Filters, error) {
	var f filter.Filter
	filters := filter.Filters{}
	rows, err := s.db.Query("SELECT * FROM filters")
	if err != nil {
		log.Fatalf("Query: %v", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&f.ID, &f.Type, &f.Value)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		filters = append(filters, f)
	}
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
