package workerprofilefetcher

import (
	"encoding/json"
	"testing"

	"github.com/gemsorg/eligibility/pkg/workerprofile"

	"github.com/gemsorg/eligibility/pkg/filter"

	"github.com/stretchr/testify/assert"
)

func TestWorkerProfileResponse(t *testing.T) {
	want := `{"id":58,"worker_id":8,"birthdate":"1982-08-05T00:00:00Z","city":"Lake Verlashire","locality":"District of Columbia","country":"Netherlands","attributes":{"Gender":[{"id":1,"type":"Gender","value":"male"}]}}`
	attr := filter.Filters{filter.Filter{1, "Gender", "male"}}
	profile := workerprofile.Profile{
		ID:         58,
		WorkerID:   8,
		Birthdate:  "1982-08-05T00:00:00Z",
		City:       "Lake Verlashire",
		Locality:   "District of Columbia",
		Country:    "Netherlands",
		Attributes: attr.GroupByType(),
	}
	actual, _ := json.Marshal(profile)
	assert.Equal(t, want, string(actual))
}
