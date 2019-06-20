package workerprofilefetcher

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/gemsorg/eligibility/pkg/authentication"

	"github.com/gemsorg/eligibility/pkg/mock"
	service "github.com/gemsorg/eligibility/pkg/service"
	"github.com/gemsorg/eligibility/pkg/workerprofile"
	"github.com/golang/mock/gomock"

	"github.com/gemsorg/eligibility/pkg/filter"

	"github.com/stretchr/testify/assert"
)

func TestWorkerProfileResponse(t *testing.T) {
	want := `{"id":58,"worker_id":8,"birthdate":"1982-08-05T00:00:00Z","city":"Lake Verlashire","locality":"District of Columbia","country":"Netherlands","state":"partial","attributes":{"Gender":[{"id":1,"type":"Gender","value":"male"}]}}`
	attr := filter.Filters{filter.Filter{1, "Gender", "male"}}
	profile := workerprofile.Profile{
		ID:         58,
		WorkerID:   8,
		Birthdate:  "1982-08-05T00:00:00Z",
		City:       "Lake Verlashire",
		Locality:   "District of Columbia",
		Country:    "Netherlands",
		State:      "partial",
		Attributes: attr.GroupByType(),
	}
	actual, _ := json.Marshal(profile)
	assert.Equal(t, want, string(actual))
}

func Test_makeCreateWorkerProfileEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockEligibilityService(ctrl)
	cxt := mock.MockContext{}
	// No error
	attr := filter.Filters{filter.Filter{1, "Gender", "male"}}
	s.EXPECT().
		SetAuthData(authentication.AuthData{}).
		AnyTimes()
	profile := workerprofile.Profile{
		ID:         58,
		WorkerID:   8,
		Birthdate:  "1982-08-05T00:00:00Z",
		City:       "Lake Verlashire",
		Locality:   "District of Columbia",
		Country:    "Netherlands",
		State:      "partial",
		Attributes: attr.GroupByType(),
	}
	s.EXPECT().
		GetWorkerProfile(gomock.Any()).
		Return(profile, nil).
		Times(1)

	resp, _ := makeWorkerProfileFetcherEndpoint(s)(cxt, WorkerProfileRequest{})
	assert.Equal(t, profile, resp)

	// Error
	err := errors.New("error creating profile")
	s.EXPECT().
		GetWorkerProfile(gomock.Any()).
		Return(workerprofile.Profile{}, err).
		Times(1)
	resp, e := makeWorkerProfileFetcherEndpoint(s)(cxt, WorkerProfileRequest{})
	assert.Equal(t, workerprofile.Profile{}, resp)
	assert.Equal(t, errorResponse(err), e)
}
