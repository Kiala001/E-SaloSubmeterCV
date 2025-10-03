package services

import (
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type CVValidationService struct {
	bus        event.Bus
	repository ports.CVRepository
}

func NewCVValidationService(repo ports.CVRepository, bus event.Bus) *CVValidationService {
	return &CVValidationService{repository: repo, bus: bus}
}

func (s *CVValidationService) ValidateCV(CvId string) error {
	cv, exists := s.repository.GetById(CvId)
	if !exists { 
		return ErrCVNotFound 
	}

	if errOrNil := cv.Validate(); errOrNil != nil { 
		return errOrNil 
	}

	s.repository.Save(cv)

	events := cv.PullEvents()

	for _, event := range events {
		s.bus.Publish(event)
	}

	return nil
}
