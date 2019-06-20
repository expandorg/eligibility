package authorization

import (
	"reflect"
	"testing"

	"github.com/gemsorg/eligibility/pkg/authentication"
)

func TestNewAuthorizer(t *testing.T) {
	tests := []struct {
		name string
		want Authorizer
	}{
		{
			"it returns an Authorizer",
			&authorizor{authentication.AuthData{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthorizer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthorizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanAccessWorkerProfile(t *testing.T) {
	type fields struct {
		authData authentication.AuthData
	}
	type args struct {
		workerID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"It returns true if authenticated user is the same as given",
			fields{authentication.AuthData{UserID: 1}},
			args{1},
			true,
			false,
		},
		{
			"It returns false if authenticated user is different from given",
			fields{authentication.AuthData{UserID: 1}},
			args{2},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizor{
				authData: tt.fields.authData,
			}
			got, err := a.CanAccessWorkerProfile(tt.args.workerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("authorizor.CanAccessWorkerProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("authorizor.CanAccessWorkerProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAuthData(t *testing.T) {
	authData := authentication.AuthData{1591960106, "http://localhost:3000", 8}
	type fields struct {
		authData authentication.AuthData
	}
	type args struct {
		data authentication.AuthData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"It sets the authdata",
			fields{authData},
			args{authData},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizor{
				authData: tt.fields.authData,
			}
			a.SetAuthData(tt.args.data)
		})
	}
}
