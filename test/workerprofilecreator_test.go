package test

// TODO: mock get profile
// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/gemsorg/eligibility/pkg/server"
// 	"github.com/gemsorg/eligibility/pkg/workerprofile"
// 	"github.com/stretchr/testify/assert"
// )

// func TestWorkerProfileCreator(t *testing.T) {
// 	db, dbx, mock := Setup()
// 	defer db.Close()
// 	wp := workerprofile.NewProfile{8, "1980-08-05", "Lake Verlashire", "District of Columbia", "Netherlands", "partial", []int{1}}
// 	// fakeProfile := workerprofile.Profile{1, 8, "1980-08-05", "Lake", "District", "Netherlands", "partial", filter.GroupedFilters{}}
// 	mock.ExpectBegin()
// 	mock.ExpectExec("REPLACE INTO worker_profiles").
// 		WithArgs(wp.WorkerID, wp.Birthdate, wp.City, wp.Locality, wp.Country, wp.State).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("Delete FROM filters_workers").WithArgs(wp.WorkerID).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("INSERT INTO filters_workers").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()
// 	mock.ExpectQuery("SELECT (.+) FROM worker_profiles WHERE worker_id=(.+) LIMIT 1").WithArgs(wp.WorkerID)
// 	tests := []struct {
// 		name   string
// 		params []byte
// 		want   int
// 	}{
// 		{
// 			"it is successful",
// 			[]byte(`{"worker_id":8,"birthdate":"1980-08-05","city":"Lake Verlashire","locality":"District of Columbia","country":"Netherlands","state":"partial","attributes":[1]}`),
// 			200,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			r, err := http.NewRequest("POST", "/workers/8/profiles", bytes.NewBuffer(tt.params))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			w := httptest.NewRecorder()
// 			s := server.New(dbx)
// 			s.ServeHTTP(w, r)

// 			resp := w.Result()
// 			// bodyBytes, _ := ioutil.ReadAll(resp.Body)
// 			// bodyString := string(bodyBytes)
// 			assert.Equal(t, tt.want, resp.StatusCode)
// 		})
// 	}
// }
