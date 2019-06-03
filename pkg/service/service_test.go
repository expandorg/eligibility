package service

import (
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/pkg/datastore"
	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ds := &datastore.EligibilityStore{}
	type args struct {
		s *datastore.EligibilityStore
	}
	tests := []struct {
		name string
		args args
		want *service
	}{
		{
			"it creates a new service",
			args{s: ds},
			&service{ds},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.s)
			assert.Equal(t, got, tt.want, tt.name)
		})
	}
}

func Test_service_Healthy(t *testing.T) {
	ds := &datastore.EligibilityStore{}
	type fields struct {
		store *datastore.EligibilityStore
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"it returns true if healthy",
			fields{store: ds},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got := s.Healthy()
			assert.Equal(t, got, tt.want, tt.name)
		})
	}
}

func Test_service_GetFilters(t *testing.T) {
	ds := &fakeStore{}
	type fields struct {
		store datastore.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		want    filter.Filters
		wantErr bool
	}{
		{
			"it gets all filters from store",
			fields{store: ds},
			fakeFilters,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetFilters()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetFilters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetFilters() = %v, want %v", got, tt.want)
			}
		})
	}
}

var fakeFilters = filter.Filters{filter.Filter{1, "Gender", "male"}}

type fakeDB struct{}

type fakeStore struct{}

func (s *fakeStore) GetAllFilters() (filter.Filters, error) {
	return fakeFilters, nil
}

func (db *fakeDB) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}
