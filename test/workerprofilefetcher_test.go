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
	r, err := http.NewRequest("GET", "/profiles/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	s := server.New(dbx)
	s.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status should be ok")
}
