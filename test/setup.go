package test

import (
	"github.com/gemsorg/eligibility/pkg/datastore"
)

func Setup() datastore.Driver {
	return &fakeDB{}
}

type fakeDB struct{}

func (db *fakeDB) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}
