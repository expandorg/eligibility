package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/expandorg/eligibility/pkg/filter"
	"github.com/expandorg/eligibility/pkg/server"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFiltersCreator(t *testing.T) {
	db, dbx, _, s := Setup(t)
	defer db.Close()
	f := filter.Filter{1, "Gender", "male"}
	tests := []struct {
		name   string
		params []byte
		want   int
	}{
		{
			"it returns true if type and value are present",
			[]byte(`{"type":"foo","value":"bar"}`),
			200,
		},
		{
			"it returns an error if request params are empty",
			[]byte(`{}`),
			500,
		},
		{
			"it returns an error if request params are empty",
			[]byte(`{"type":"","value":""}`),
			500,
		},
		{
			"it returns an error if type is empty",
			[]byte(`{"type":"", "value": "bar"}`),
			500,
		},
		{
			"it returns an error if value is empty",
			[]byte(`{"type":"foo", "value": ""}`),
			500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.EXPECT().
				CreateFilter(gomock.Any()).
				Return(f, nil).
				Times(1)

			s.EXPECT().
				SetAuthData(gomock.Any()).
				AnyTimes()
			r, err := http.NewRequest("POST", "/filters", bytes.NewBuffer(tt.params))
			if err != nil {
				t.Fatal(err)
			}
			r.Header.Add("Authorization", bearer)
			w := httptest.NewRecorder()
			svr := server.New(dbx, s)
			svr.ServeHTTP(w, r)

			resp := w.Result()
			assert.Equal(t, tt.want, resp.StatusCode)
		})
	}
}
