package mock

import (
	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
)

var FakeFilter = filter.Filter{1, "Gender", "male"}
var FakeFilters = filter.Filters{FakeFilter}
var EmptyProfile = workerprofile.Profile{Attributes: filter.GroupedFilters{}}
var FakeProfile = workerprofile.Profile{
	ID:         58,
	WorkerID:   8,
	Birthdate:  "1982-08-05T00:00:00Z",
	City:       "Lake Verlashire",
	Locality:   "District of Columbia",
	Country:    "Netherlands",
	Attributes: FakeFilters.GroupByType(),
}

type FakeDB struct{}

type FakeStore struct{}

func (s *FakeStore) GetAllFilters(ids []int) (filter.Filters, error) {
	return FakeFilters, nil
}

func (s *FakeStore) CreateFilter(filter.Filter) (filter.Filter, error) {
	return FakeFilter, nil
}

func (s *FakeStore) GetWorkerProfile(workerID string) (workerprofile.Profile, error) {
	if workerID == "nonExistantID" {
		return EmptyProfile, nil
	}
	return FakeProfile, nil
}

func (s *FakeStore) CreateWorkerProfile(wp workerprofile.NewProfile) (workerprofile.Profile, error) {
	return FakeProfile, nil
}

func (db *FakeDB) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}
