package datastore

import (
	"testing"

	"github.com/gemsorg/eligibility/pkg/filter"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gemsorg/eligibility/mock"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
	"github.com/jmoiron/sqlx"
)

func TestEligibilityStore_CreateWorkerProfile(t *testing.T) {
	_, dbx, mock := mock.Mysql()
	defer dbx.Close()
	wp := workerprofile.NewProfile{8, "1980-08-05", "Lake", "District", "Netherlands", []int{}}
	fakeProfile := workerprofile.Profile{1, 8, "1980-08-05", "Lake", "District", "Netherlands", filter.GroupedFilters{}}
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		wp workerprofile.NewProfile
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   workerprofile.Profile
	}{
		{
			"it runs everything in transaction",
			fields{dbx},
			args{wp},
			fakeProfile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EligibilityStore{
				DB: tt.fields.DB,
			}
			mock.ExpectBegin()
			mock.ExpectExec("REPLACE INTO worker_profiles").
				WithArgs(wp.WorkerID, wp.Birthdate, wp.City, wp.Locality, wp.Country).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec("Delete FROM filters_workers").WithArgs(wp.WorkerID).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec("INSERT INTO filters_workers").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			s.CreateWorkerProfile(tt.args.wp)
			// p, err := s.CreateWorkerProfile(tt.args.wp)
			// TODO: make these assertions work
			// assert.Equal(t, p, fakeProfile)
			// assert.Nil(t, err)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})
	}
}
