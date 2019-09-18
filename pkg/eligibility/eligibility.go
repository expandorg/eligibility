package eligibility

type WorkerEligibility struct {
	Complete    bool     `json:"complete"`
	Eligibile   []uint64 `json:"eligible"`
	InEligibile []uint64 `json:"ineligible"`
}
