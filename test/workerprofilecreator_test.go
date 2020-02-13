package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/expandorg/eligibility/pkg/authentication"
	"github.com/expandorg/eligibility/pkg/filter"
	"github.com/expandorg/eligibility/pkg/server"
	"github.com/expandorg/eligibility/pkg/workerprofile"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWorkerProfileCreator(t *testing.T) {
	db, dbx, _, s := Setup(t)
	defer db.Close()

	attr := filter.Filters{filter.Filter{1, "Gender", "male"}}
	profile := workerprofile.Profile{
		ID:         58,
		WorkerID:   8,
		Name:       "Test User",
		Birthdate:  "1982-08-05T00:00:00Z",
		City:       "Lake Verlashire",
		Locality:   "District of Columbia",
		Country:    "Netherlands",
		State:      "partial",
		Attributes: attr.GroupByType(),
	}
	s.EXPECT().
		CreateWorkerProfile(gomock.Any()).
		Return(profile, nil).
		Times(1)

	s.EXPECT().
		SetAuthData(gomock.Any()).
		AnyTimes()

	s.EXPECT().
		GetWorkerProfile(gomock.Any()).
		Return(profile, nil).
		AnyTimes()
	s.SetAuthData(authentication.AuthData{UserID: 8})
	tests := []struct {
		name   string
		params []byte
		want   int
	}{
		{
			"it is successful",
			[]byte(`{"worker_id":8,"name":"Test User","birthdate":"1980-08-05","city":"Lake Verlashire","locality":"District of Columbia","country":"Netherlands","state":"partial","attributes":[1]}`),
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r, err := http.NewRequest("POST", "/workers/8/profiles", bytes.NewBuffer(tt.params))
			r.Header.Add("Authorization", bearer)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			srvr := server.New(dbx, s)
			srvr.ServeHTTP(w, r)

			resp := w.Result()
			assert.Equal(t, tt.want, resp.StatusCode)
			var p workerprofile.Profile
			decoder := json.NewDecoder(resp.Body)
			decoder.Decode(&p)
			assert.Equal(t, profile, p)
		})
	}
}
