package service

import (
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/mock"
	"github.com/gemsorg/eligibility/pkg/datastore"
	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
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
	ds := &mock.FakeStore{}
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
			mock.FakeFilters,
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

func Test_service_CreateFilter(t *testing.T) {
	ds := &mock.FakeStore{}
	requestFilter := mock.FakeFilter
	requestFilter.ID = 0
	type fields struct {
		store datastore.Storage
	}
	type args struct {
		f filter.Filter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    filter.Filter
		wantErr bool
	}{
		{
			"it creates a new filter",
			fields{ds},
			args{requestFilter},
			mock.FakeFilter,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.CreateFilter(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetWorkerProfile(t *testing.T) {
	ds := &mock.FakeStore{}
	type fields struct {
		store datastore.Storage
	}
	type args struct {
		workerID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    workerprofile.Profile
		wantErr bool
	}{
		{
			"it returns the worker's profile",
			fields{ds},
			args{"1"},
			mock.FakeProfile,
			false,
		},
		{
			"it returns an empty profile if worker_id not found",
			fields{ds},
			args{"nonExistantID"},
			mock.EmptyProfile,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetWorkerProfile(tt.args.workerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetWorkerProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetWorkerProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
