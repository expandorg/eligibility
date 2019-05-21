package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gemsorg/eligibility/log"
)

func TestNew(t *testing.T) {
	l := log.New()
	type args struct {
		l log.Logger
	}
	tests := []struct {
		name string
		args args
		want *service
	}{
		{
			"it returns a service with logger",
			args{l},
			&service{l},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.l)
			assert.Equal(t, got, tt.want, tt.name)
		})
	}
}

func Test_service_Healthy(t *testing.T) {
	type fields struct {
		logger log.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"it returns true",
			fields{log.New()},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger: tt.fields.logger,
			}
			got := s.Healthy()
			assert.Equal(t, got, tt.want, tt.name)
		})
	}
}
