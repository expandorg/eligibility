package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemsorg/eligibility/pkg/filter"
	"github.com/gemsorg/eligibility/pkg/server"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFiltersFetcher(t *testing.T) {
	db, dbx, _, s := Setup(t)
	defer db.Close()
	fs := filter.Filters{filter.Filter{1, "Gender", "male"}}
	s.EXPECT().
		GetFilters().
		Return(fs, nil).
		Times(1)

	s.EXPECT().
		SetAuthData(gomock.Any()).
		AnyTimes()
	r, err := http.NewRequest("GET", "/filters", nil)
	r.Header.Add("Authorization", bearer)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	svr := server.New(dbx, s)
	svr.ServeHTTP(w, r)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status should be ok")
	var f filter.GroupedFilters
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&f)
	assert.Equal(t, fs.GroupByType(), f)
}
