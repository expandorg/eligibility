package eligibility

import (
	"reflect"
	"testing"

	"github.com/expandorg/eligibility/pkg/filter"
)

func TestGetWorkerEligibility(t *testing.T) {
	type args struct {
		wf filter.FilterWorker
		js []filter.FilterJob
		c  bool
	}
	usa := uint64(1)
	spain := uint64(2)
	italy := uint64(3)

	tests := []struct {
		name string
		args args
		want WorkerEligibility
	}{
		{
			"equal comparison: eligible for usa job only",
			args{
				filter.FilterWorker{1, usa},
				[]filter.FilterJob{
					filter.FilterJob{1, usa, "=="},
					filter.FilterJob{2, italy, "=="},
					filter.FilterJob{3, spain, "=="},
				},
				true,
			},
			WorkerEligibility{Complete: true, Eligible: []uint64{usa}, InEligible: []uint64{2, 3}},
		},
		{
			"not equal comparison: eligible for italy and spain jobs only",
			args{
				filter.FilterWorker{1, usa},
				[]filter.FilterJob{
					filter.FilterJob{1, italy, "!="},
					filter.FilterJob{2, spain, "!="},
					filter.FilterJob{3, usa, "!="},
				},
				true,
			},
			WorkerEligibility{Complete: true, Eligible: []uint64{1, 2}, InEligible: []uint64{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWorkerEligibility(tt.args.wf, tt.args.js, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkerEligibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
