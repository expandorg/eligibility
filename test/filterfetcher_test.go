package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemsorg/eligibility/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestFiltersFetcher(t *testing.T) {
	_, dbx, _ := Setup()
	r, err := http.NewRequest("GET", "/filters", nil)
	r.Header.Add("Authorization", bearer)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	s := server.New(dbx)
	s.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status should be ok")
}
