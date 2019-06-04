package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gemsorg/eligibility/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestFiltersCreator(t *testing.T) {
	db, dbx, mock := Setup()
	defer db.Close()
	mock.ExpectExec("INSERT INTO filters").WithArgs("foo", "bar").
		WillReturnResult(sqlmock.NewResult(0, 0))
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

			r, err := http.NewRequest("POST", "/filters", bytes.NewBuffer(tt.params))
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			s := server.New(dbx)
			s.ServeHTTP(w, r)

			resp := w.Result()
			// bodyBytes, _ := ioutil.ReadAll(resp.Body)
			// bodyString := string(bodyBytes)
			assert.Equal(t, tt.want, resp.StatusCode)
		})
	}
}
