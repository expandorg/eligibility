package workerprofilecreator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/pkg/apierror"
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
