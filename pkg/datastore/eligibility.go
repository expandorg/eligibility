package datastore

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/expandorg/eligibility/pkg/filter"
	"github.com/expandorg/eligibility/pkg/workerprofile"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetAllFilters(ids []int) (filter.Filters, error)
	CreateFilter(filter.Filter) (filter.Filter, error)
	GetWorkerProfile(workerID string) (workerprofile.Profile, error)
	CreateWorkerProfile(wp workerprofile.NewProfile) (workerprofile.Profile, error)
	GetWorkerEligibility(workerID string) (filter.FilterWorker, []filter.FilterJob, error)
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
	p := workerprofile.Profile{State: workerprofile.NOTFILLED, Attributes: attr}

	err := s.DB.Get(&p, "SELECT * FROM worker_profiles WHERE worker_id=? LIMIT 1", workerID)
	if err != nil {
		if err == sql.ErrNoRows {
			authUserID, _ := strconv.ParseUint(workerID, 10, 64)
			p.ID = authUserID
			return p, nil
		}
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

	// Create new location filters if they don't exist
	locIDS, err := s.locationToFilters(tx, wp)

	if err != nil {
		tx.Rollback()
		return workerprofile.Profile{}, err
	}

	wp.Attributes = append(wp.Attributes, locIDS...)

	_, err = tx.Exec(
		"REPLACE INTO worker_profiles (worker_id, name, birthdate, city, locality, country, state) VALUES (?, ?, ?, ?, ?, ?, ?)",
		wp.WorkerID, wp.Name, wp.Birthdate, wp.City, wp.Locality, wp.Country, wp.State)

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

// We convert (find or create) location profile data (city, locality, country) to filters
func (s *EligibilityStore) locationToFilters(tx *sql.Tx, wp workerprofile.NewProfile) ([]int, error) {
	locIDS := []int{}
	a := map[string]string{
		"Country":  wp.Country,
		"City":     wp.City,
		"Locality": wp.Locality,
	}

	type filterResponse struct {
		ID  int64
		Err error
	}
	ch := make(chan filterResponse)

	for t, v := range a {
		if v == "" {
			delete(a, t)
			continue
		}
		go func(tp string, value string) {
			id, err := s.findOrCreateFilter(tx, tp, value)
			if err != nil {
				tx.Rollback()
				ch <- filterResponse{Err: err}
			}
			ch <- filterResponse{ID: id}
		}(t, v)
	}

	working := true
	for len(a) > 0 && working {
		select {
		case r := <-ch:
			if r.Err != nil {
				tx.Rollback()
				return []int{}, r.Err
			}
			locIDS = append(locIDS, int(r.ID))
			// keep listening/waiting until we're done with all filters
			if len(locIDS) == len(a) {
				close(ch)
				working = false
			}
		}
	}
	return locIDS, nil
}

func (s *EligibilityStore) findOrCreateFilter(tx *sql.Tx, tp string, value string) (int64, error) {
	res, err := tx.Exec(
		"INSERT INTO filters (type, value) VALUES (?, ?) ON DUPLICATE KEY UPDATE value=?",
		tp, value, value,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var id int64
	rows, _ := res.RowsAffected()
	// If a filter already existed, it won't be returned from the above query, so we select it
	if rows == 0 {
		err = s.DB.Get(&id, "SELECT id FROM filters WHERE type=? AND value=?", tp, value)
		if err != nil {
			return 0, err
		}
	} else {
		// we created a new filter, get the inserted id
		id, _ = res.LastInsertId()
	}
	return id, nil
}

func (s *EligibilityStore) GetWorkerEligibility(workerID string) (filter.FilterWorker, []filter.FilterJob, error) {
	fw := filter.FilterWorker{}
	fj := []filter.FilterJob{}

	err := s.DB.Get(&fw, "SELECT fj.`worker_id`, fj.`filter_id` from filters AS f inner join filters_workers AS fj on f.`id` = fj.`filter_id` where fj.`worker_id` = ? AND f.type =?", workerID, "Country")
	if err != nil && err != sql.ErrNoRows {
		return fw, fj, err
	}

	err = s.DB.Select(&fj, "SELECT fj.`job_id`, fj.`filter_id`, fj.`comparison` from filters AS f inner join filters_jobs AS fj on f.`id` = fj.`filter_id` where f.`type`=?", "Country")
	if err != nil {
		if err == sql.ErrNoRows {
			return fw, fj, nil
		}
		return fw, fj, err
	}

	return fw, fj, nil
}
