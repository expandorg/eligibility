package datastore

import (
	"errors"
	"time"
)

func (s *EligibilityStore) validateFiltersExist(ids []int) (bool, error) {
	fs, err := s.GetAllFilters(ids)

	if err != nil {
		return false, err
	}

	if len(ids) > len(fs) {
		filterNotFound := FilterNotFound{ids}
		return false, errors.New(filterNotFound.Error())
	}
	return true, nil
}

func validateDateFormat(s string) (time.Time, error) {

	t, err := time.Parse("2006-01-02", s)

	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
