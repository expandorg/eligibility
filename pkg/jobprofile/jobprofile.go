package jobprofile

import "github.com/expandorg/eligibility/pkg/filter"

type JobProfile struct {
	Countries  []string              `json:"countries"`
	Cities     []string              `json:"cities"`
	Locality   []string              `json:"locality"`
	State      []string              `json:"state"`
	Attributes filter.GroupedFilters `json:"attributes"`
}
