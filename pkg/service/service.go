package service

import (
	"github.com/gemsorg/eligibility/pkg/datastore"

	"github.com/gemsorg/eligibility/pkg/filter"
)

type EligibilityService interface {
	Healthy() bool
	GetFilters() (filter.Filters, error)
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
