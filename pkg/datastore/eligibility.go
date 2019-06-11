package datastore

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetAllFilters(ids []int) (filter.Filters, error)
	CreateFilter(filter.Filter) (filter.Filter, error)
	GetWorkerProfile(workerID string) (workerprofile.Profile, error)
	CreateWorkerProfile(wp workerprofile.NewProfile) (workerprofile.Profile, error)
}

type EligibilityStore struct {
	DB *sqlx.DB
}

func NewEligibilityStore(db *sqlx.DB) *EligibilityStore {
	return &EligibilityStore{
		DB: db,
	}
}

func (s *EligibilityStore) GetAllFilters(ids []int) (filter.Filters, error) {
	filters := filter.Filters{}
	if len(ids) > 0 {
		s.DB.Select(&filters, "SELECT * FROM filters WHERE id in(?)", ids)
	} else {
		s.DB.Select(&filters, "SELECT * FROM filters")
	}
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

func (s *EligibilityStore) CreateWorkerProfile(wp workerprofile.NewProfile) (workerprofile.Profile, error) {
	tx, err := s.DB.Begin()
	_, err = tx.Exec(
		"REPLACE INTO worker_profiles (worker_id, birthdate, city, locality, country) VALUES (?, ?, ?, ?, ?)",
		wp.WorkerID, wp.Birthdate, wp.City, wp.Locality, wp.Country)

	if err != nil {
		tx.Rollback()
		return workerprofile.Profile{}, err
	}

	_, err = tx.Exec("Delete FROM filters_workers WHERE worker_id=?", wp.WorkerID)

	if err != nil {
		tx.Rollback()
		return workerprofile.Profile{}, err
	}

	if len(wp.Attributes) > 0 {
		vals := []string{}

		for _, id := range wp.Attributes {
			vals = append(vals, fmt.Sprintf("(%d, %d)", wp.WorkerID, id))
		}
		attrQuery := "INSERT INTO filters_workers (worker_id, filter_id) VALUES" + strings.Join(vals, ",")
		_, err = tx.Exec(attrQuery)

		if err != nil {
			tx.Rollback()
			return workerprofile.Profile{}, FilterNotFound{wp.Attributes}
		}
	}

	err = tx.Commit()

	if err != nil {
		return workerprofile.Profile{}, err
	}

	p, err := s.GetWorkerProfile(strconv.FormatUint(wp.WorkerID, 10))

	if err != nil {
		return workerprofile.Profile{}, err
	}
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
