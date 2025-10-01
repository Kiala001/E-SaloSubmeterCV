package services

import (
	"errors"
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
	CV, exists := s.repository.GetById(CvId)
	if !exists { 
		return errors.New("CV not found") 
	}

	if errOrNil := CV.Validate(); errOrNil != nil { 
		return errOrNil 
	}

	s.repository.Save(CV)

	events := CV.PullEvents()
	CV.PublishEvent(s.bus, events[0])

	return nil
}
