package filtercreator

import (
	"reflect"
	"testing"

	"github.com/expandorg/eligibility/pkg/apierror"
	"github.com/expandorg/eligibility/pkg/filter"
)

func Test_validateRequest(t *testing.T) {
	type args struct {
		req filter.Filter
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 *apierror.APIError
	}{
		{
			"is valid if type and value are not empty",
			args{filter.Filter{0, "foo", "bar"}},
			true,
			nil,
		},
		{
			"is not valid if type is empty",
			args{filter.Filter{0, "", "foo"}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"type"}}),
		},
		{
			"is not valid if value is empty",
			args{filter.Filter{0, "foo", ""}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"value"}}),
		},
		{
			"is not valid if value and type are empty",
			args{filter.Filter{0, "", ""}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"type", "value"}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateRequest(tt.args.req)
			if got != tt.want {
				t.Errorf("validateRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
