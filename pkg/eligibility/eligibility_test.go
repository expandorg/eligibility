package eligibility

import (
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/pkg/filter"
)

func TestGetWorkerEligibility(t *testing.T) {
	type args struct {
		wf filter.FilterWorker
		js []filter.FilterJob
		c bool
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
			"it returns true",
			args{
				filter.FilterWorker{1, usa},
				[]filter.FilterJob{
					filter.FilterJob{1, usa, "=="},
					filter.FilterJob{1, italy, "=="},
					filter.FilterJob{2, spain, "=="},
				},
				true,
			},
			WorkerEligibility{Complete: true, Eligibile: []uint64{usa}, InEligibile: []uint64{2}},
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

func TestGetWorkerEligibilityForNotEqualsComparison(t *testing.T) {
	type args struct {
		wf filter.FilterWorker
		js []filter.FilterJob
		c bool
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
			"it returns true",
			args{
				filter.FilterWorker{1, usa},
				[]filter.FilterJob{
					filter.FilterJob{1, italy, "!="},
					filter.FilterJob{2, spain, "!="},
					filter.FilterJob{3, usa, "!="},
				},
				true,
			},
			WorkerEligibility{Complete: true, Eligibile: []uint64{1, 2}, InEligibile: []uint64{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWorkerEligibilityForNotEqualsComparison(tt.args.wf, tt.args.js, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkerEligibilityForNotEqualsComparison() = %v, want %v", got, tt.want)
			}
		})
	}
}