package filter

import (
	"reflect"
	"testing"
)

func TestFilters_GroupByType(t *testing.T) {
	tests := []struct {
		name string
		fs   *Filters
		want GroupedFilters
	}{
		{
			"It groups filters by type",
			&Filters{
				{1, "Gender", "male"},
				{2, "Education", "bachelor"},
				{3, "Gender", "female"},
				{4, "Education", "masters"},
			},
			GroupedFilters{
				"Gender":    Filters{{1, "Gender", "male"}, {3, "Gender", "female"}},
				"Education": Filters{{2, "Education", "bachelor"}, {4, "Education", "masters"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fs.GroupByType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filters.GroupByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
