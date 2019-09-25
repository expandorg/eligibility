package datastore

// func TestEligibilityStore_CreateWorkerProfile(t *testing.T) {
// 	_, dbx, mock := mock.Mysql()
// 	defer dbx.Close()
// 	wp := workerprofile.NewProfile{8, "Test User", "1980-08-05", "Lake", "District", "Netherlands", "partial", []int{1}}
// 	fakeProfile := workerprofile.Profile{1, 8, "Test User", "1980-08-05", "Lake", "District", "Netherlands", "partial", filter.GroupedFilters{}}
// 	type fields struct {
// 		DB *sqlx.DB
// 	}
// 	type args struct {
// 		wp workerprofile.NewProfile
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   workerprofile.Profile
// 	}{
// 		{
// 			"it runs everything in transaction",
// 			fields{dbx},
// 			args{wp},
// 			fakeProfile,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			smock := NewMockStorage(ctrl)

// 			// No error
// 			smock.EXPECT().
// 				LocationToFilters(dbx, wp)

// 			mock.ExpectBegin()
// 			mock.ExpectExec("REPLACE INTO worker_profiles").
// 				WithArgs(wp.WorkerID, wp.Name, wp.Birthdate, wp.City, wp.Locality, wp.Country, wp.State).
// 				WillReturnResult(sqlmock.NewResult(1, 1))
// 			mock.ExpectExec("Delete FROM filters_workers").WithArgs(wp.WorkerID).WillReturnResult(sqlmock.NewResult(1, 1))
// 			mock.ExpectExec("INSERT INTO filters_workers").WillReturnResult(sqlmock.NewResult(1, 1))
// 			mock.ExpectCommit()

// 			smock.CreateWorkerProfile(tt.args.wp)
// 			// p, err := s.CreateWorkerProfile(tt.args.wp)
// 			// TODO: make these assertions work
// 			// assert.Equal(t, p, fakeProfile)
// 			// assert.Nil(t, err)
// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}

// 		})
// 	}
// }
