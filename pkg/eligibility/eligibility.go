package eligibility

import (
	"log"

	"github.com/gemsorg/eligibility/pkg/filter"
)

type WorkerEligibility struct {
	Complete    bool     `json:"complete"`
	Eligibile   []uint64 `json:"eligible"`
	InEligibile []uint64 `json:"ineligible"`
}

// For now, we're only supporting filtering by country
func GetWorkerEligibility(wf filter.FilterWorker, js []filter.FilterJob, profileComplete bool) WorkerEligibility {
	we := WorkerEligibility{
		Complete:    profileComplete,
		Eligibile:   []uint64{},
		InEligibile: []uint64{},
	}
	el := map[uint64]bool{}
	iel := map[uint64]bool{}

	for _, j := range js {

		switch j.Comparison {
		case "==":
			if j.FilterID == wf.FilterID {
				we.Eligibile = append(we.Eligibile, j.JobID)
				el[j.JobID] = true
			} else {
				we.InEligibile = append(we.InEligibile, j.JobID)
			}
		case "!=":
			if j.FilterID == wf.FilterID {
				we.InEligibile = append(we.InEligibile, j.JobID)
				iel[j.JobID] = true
			} else {
				we.Eligibile = append(we.Eligibile, j.JobID)
			}
		default:
			continue
		}
	}

	// if a job is already eligible for a country, remove it from ineligible list.
	for i, inel := range we.InEligibile {
		if el[inel] {
			we.InEligibile = append(we.InEligibile[:i], we.InEligibile[i+1:]...)
		}
	}

	// if a job is already ineligible for a country, remove it from eligible list.
	for i, inel := range we.Eligibile {
		if iel[inel] {
			we.Eligibile = append(we.Eligibile[:i], we.Eligibile[i+1:]...)
		}
	}

	log.Fatal(we)
	return we
}
