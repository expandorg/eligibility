package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemsorg/eligibility/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestWorkerProfileFetcher(t *testing.T) {
	_, dbx, _ := Setup()
	r, err := http.NewRequest("GET", "/workers/1/profiles", nil)
	r.Header.Add("Authorization", bearer)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	s := server.New(dbx)
	s.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status should be ok")
}
