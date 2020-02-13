package filtercreator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/expandorg/eligibility/pkg/apierror"
	"github.com/expandorg/eligibility/pkg/filter"
	"github.com/expandorg/eligibility/pkg/mock"
	service "github.com/expandorg/eligibility/pkg/service"
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

func Test_makeCreateFilterEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockEligibilityService(ctrl)
	cxt := mock.MockContext{}

	// No error
	f := filter.Filter{1, "Gender", "male"}
	s.EXPECT().
		CreateFilter(f).
		Return(f, nil).
		Times(1)

	resp, _ := makeCreateFilterEndpoint(s)(cxt, f)
	assert.Equal(t, f, resp)

	// Error
	err := errors.New("error creating filter")
	s.EXPECT().
		CreateFilter(f).
		Return(filter.Filter{}, err).
		Times(1)
	resp, e := makeCreateFilterEndpoint(s)(cxt, f)
	assert.Equal(t, nil, resp)
	assert.Equal(t, errorResponse(err), e)
}
