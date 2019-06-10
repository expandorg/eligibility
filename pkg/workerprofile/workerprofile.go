package workerprofile

import (
	"github.com/gemsorg/eligibility/pkg/filter"
)

type Profile struct {
	ID         uint64                `json:"id"`
	WorkerID   uint64                `json:"worker_id" db:"worker_id"`
	Birthdate  string                `json:"birthdate"`
	City       string                `json:"city"`
	Locality   string                `json:"locality"`
	Country    string                `json:"country"`
	Attributes filter.GroupedFilters `json:"attributes"`
}

type NewProfile struct {
	WorkerID   uint64 `json:"worker_id"`
	Birthdate  string `json:"birthdate"`
	City       string `json:"city"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
	Attributes []int  `json:"attributes"`
}
