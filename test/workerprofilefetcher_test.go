package test

import (
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

func TestWorkerProfileFetcher(t *testing.T) {
	_, dbx, _, s := Setup(t)

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
		GetWorkerProfile(gomock.Any()).
		Return(profile, nil).
		AnyTimes()

	s.EXPECT().
		SetAuthData(gomock.Any()).
		AnyTimes()

	s.SetAuthData(authentication.AuthData{UserID: 8})

	w := httptest.NewRecorder()
	sv := server.New(dbx, s)
	r, err := http.NewRequest("GET", "/workers/8/profiles", nil)
	r.Header.Add("Authorization", bearer)
	if err != nil {
		t.Fatal(err)
	}
	sv.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status should be ok")
	var p workerprofile.Profile
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&p)
	assert.Equal(t, profile, p)
}
