package service

import (
	"fmt"

	"github.com/gemsorg/eligibility/log"
)

type EligibilityService interface {
	Healthy() bool
}

type service struct {
	logger log.Logger
}

func New(l log.Logger) *service {
	return &service{
		logger: l,
	}
}

func (s *service) Healthy() bool {
	s.logger.Log("health", fmt.Sprintf("I'm healthy!"))
	return true
}
