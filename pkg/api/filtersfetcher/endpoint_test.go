package filtersfetcher

import (
	"encoding/json"
	"testing"

	"github.com/gemsorg/eligibility/pkg/filter"

	"github.com/stretchr/testify/assert"
)

func TestFiltersResponse(t *testing.T) {
	want := `{"Gender":[{"id":1,"type":"Gender","value":"male"}]}`
	fitlers := filter.Filters{filter.Filter{1, "Gender", "male"}}
	grouped := fitlers.GroupByType()
	actual, _ := json.Marshal(grouped)
	assert.Equal(t, want, string(actual))
}
