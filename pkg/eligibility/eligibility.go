package eligibility

import (
	"github.com/gemsorg/eligibility/pkg/filter"
)

type WorkerEligibility struct {
	Complete    bool     `json:"complete"`
	Eligible   []uint64 `json:"eligible"`
	InEligible []uint64 `json:"ineligible"`
}

// For now, we're only supporting filtering by country
func GetWorkerEligibility(wf filter.FilterWorker, js []filter.FilterJob, profileComplete bool) WorkerEligibility {
	we := WorkerEligibility{
		Complete:    profileComplete,
		Eligible:   []uint64{},
		InEligible: []uint64{},
	}
	el := map[uint64]bool{}
	iel := map[uint64]bool{}

	for _, j := range js {

		switch j.Comparison {
		case "==":
			if j.FilterID == wf.FilterID {
				we.Eligible = append(we.Eligible, j.JobID)
				el[j.JobID] = true
			} else {
				we.InEligible = append(we.InEligible, j.JobID)
			}
		case "!=":
			if j.FilterID == wf.FilterID {
				we.InEligible = append(we.InEligible, j.JobID)
				iel[j.JobID] = true
			} else {
				we.Eligible = append(we.Eligible, j.JobID)
			}
		default:
			continue
		}
	}

	// if a job is already eligible for a country, remove it from ineligible list.
	for i, inel := range we.InEligible {
		if el[inel] {
			we.InEligible = append(we.InEligible[:i], we.InEligible[i+1:]...)
		}
	}

	// if a job is already ineligible for a country, remove it from eligible list.
	for i, inel := range we.Eligible {
		if iel[inel] {
			we.Eligible = append(we.Eligible[:i], we.Eligible[i+1:]...)
		}
	}

	return we
}
