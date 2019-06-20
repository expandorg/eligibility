package filtersfetcher

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/pkg/apierror"
	"github.com/gemsorg/eligibility/pkg/filter"
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

func Test_makeFiltersFetcherEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockEligibilityService(ctrl)
	cxt := mock.MockContext{}

	// No error
	filters := filter.Filters{
		{1, "Gender", "male"},
		{2, "Education", "bachelor"},
		{3, "Gender", "female"},
		{4, "Education", "masters"},
	}
	grouped := filter.GroupedFilters{
		"Gender":    filter.Filters{{1, "Gender", "male"}, {3, "Gender", "female"}},
		"Education": filter.Filters{{2, "Education", "bachelor"}, {4, "Education", "masters"}},
	}

	s.
		EXPECT().
		GetFilters().
		Return(filters, nil).
		Times(1)
	resp, _ := makeFiltersFetcherEndpoint(s)(cxt, "request")
	assert.Equal(t, resp, grouped)

	// Error
	err := errors.New("error fetching filters")
	s.
		EXPECT().
		GetFilters().
		Return(filter.Filters{}, err).
		Times(1)

	resp, e := makeFiltersFetcherEndpoint(s)(cxt, "request")
	assert.Equal(t, resp, nil)
	assert.Equal(t, errorResponse(err), e)
}
