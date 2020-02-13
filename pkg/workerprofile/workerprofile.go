package workerprofile

import (
	"github.com/expandorg/eligibility/pkg/filter"
)

const (
	NOTFILLED  = "not_filled"
	INCOMPLETE = "incomplete"
	COMPLETE   = "complete"
)

type Profile struct {
	ID         uint64                `json:"id"`
	WorkerID   uint64                `json:"worker_id" db:"worker_id"`
	Name       string                `json:"name"`
	Birthdate  string                `json:"birthdate"`
	City       string                `json:"city"`
	Locality   string                `json:"locality"`
	Country    string                `json:"country"`
	State      string                `json:"state"`
	Attributes filter.GroupedFilters `json:"attributes"`
}

type NewProfile struct {
	WorkerID   uint64 `json:"worker_id"`
	Name       string `json:"name"`
	Birthdate  string `json:"birthdate"`
	City       string `json:"city"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
	State      string `json:"state"`
	Attributes []int  `json:"attributes"`
}
