package workerprofilecreator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/expandorg/eligibility/pkg/authentication"
	"github.com/expandorg/eligibility/pkg/filter"
	"github.com/expandorg/eligibility/pkg/mock"
	service "github.com/expandorg/eligibility/pkg/service"
	"github.com/expandorg/eligibility/pkg/workerprofile"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/expandorg/eligibility/pkg/apierror"
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

func Test_makeCreateWorkerProfileEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockEligibilityService(ctrl)
	cxt := mock.MockContext{}
	newProfile := workerprofile.NewProfile{}
	// No error
	attr := filter.Filters{filter.Filter{1, "Gender", "male"}}
	profile := workerprofile.Profile{
		ID:         58,
		WorkerID:   8,
		Name:       "Test User",
		Birthdate:  "1982-08-05T00:00:00Z",
		City:       "Lake Verlashire",
		Locality:   "District of Columbia",
		Country:    "Netherlands",
		State:      "partial",
		Attributes: attr.GroupByType(),
	}
	s.EXPECT().
		SetAuthData(authentication.AuthData{}).
		AnyTimes()

	s.EXPECT().
		CreateWorkerProfile(newProfile).
		Return(profile, nil).
		Times(1)

	resp, _ := makeCreateWorkerProfileEndpoint(s)(cxt, newProfile)
	assert.Equal(t, profile, resp)

	// Error
	err := errors.New("error creating profile")
	s.EXPECT().
		CreateWorkerProfile(newProfile).
		Return(workerprofile.Profile{}, err).
		Times(1)
	resp, e := makeCreateWorkerProfileEndpoint(s)(cxt, newProfile)
	assert.Equal(t, nil, resp)
	assert.Equal(t, errorResponse(err), e)
}
