package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/expandorg/eligibility/pkg/eligibility"
	"github.com/expandorg/eligibility/pkg/server"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEligibilityFetcher(t *testing.T) {
	_, dbx, _, s := Setup(t)

	el := eligibility.WorkerEligibility{
		true,
		[]uint64{1},
		[]uint64{2},
	}
	s.EXPECT().
		GetWorkerEligibility(gomock.Any()).
		Return(el, nil).
		AnyTimes()

	w := httptest.NewRecorder()
	sv := server.New(dbx, s)
	r, err := http.NewRequest("GET", "/workers/1/eligibility", nil)
	r.Header.Add("Authorization", bearer)
	if err != nil {
		t.Fatal(err)
	}
	sv.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status should be ok")
	var e eligibility.WorkerEligibility
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&e)
	assert.Equal(t, el, e)
}
