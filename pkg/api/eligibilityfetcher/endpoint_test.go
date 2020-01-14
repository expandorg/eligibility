package eligibilityfetcher

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/pkg/apierror"
	"github.com/gemsorg/eligibility/pkg/eligibility"
	"github.com/gemsorg/eligibility/pkg/mock"
	"github.com/gemsorg/eligibility/pkg/service"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func Test_errorResponse(t *testing.T) {
	msg := "error message"
	err := errors.New(msg)
	apiErr := apierror.New(500, msg, err)
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *apierror.APIError
	}{
		{
			"it returns an API error",
			args{err},
			apiErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorResponse(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeEligibilityFetcherEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockEligibilityService(ctrl)
	cxt := mock.MockContext{}

	// No error
	el := eligibility.WorkerEligibility{
		Complete:    false,
		Eligible:   []uint64{},
		InEligible: []uint64{1},
	}

	s.
		EXPECT().
		GetWorkerEligibility(gomock.Any()).
		Return(el, nil).
		Times(1)

	resp, _ := makeEligibilityFetcherEndpoint(s)(cxt, WorkerEligibilityRequest{})
	assert.Equal(t, resp, el)

	// Error
	err := errors.New("error fetching filters")
	s.
		EXPECT().
		GetWorkerEligibility(gomock.Any()).
		Return(eligibility.WorkerEligibility{}, err).
		Times(1)

	resp, e := makeEligibilityFetcherEndpoint(s)(cxt, WorkerEligibilityRequest{})
	assert.Equal(t, resp, nil)
	assert.Equal(t, errorResponse(err), e)
}
