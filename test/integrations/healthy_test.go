package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemsorg/eligibility/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	r, err := http.NewRequest("GET", "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	s := server.New()
	s.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status should be ok")
}
