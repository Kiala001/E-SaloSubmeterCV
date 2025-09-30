package services

import (
	"errors"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type CVValidationService struct {
	Bus        event.Bus
	Repository ports.CVRepository
}

func NewCVValidationService(repo ports.CVRepository, bus event.Bus) *CVValidationService {
	return &CVValidationService{Repository: repo, Bus: bus}
}

func (s *CVValidationService) ValidateCV(CvId string) error {
	CV, exists := s.Repository.GetById(CvId)
	if !exists { return errors.New("CV not found") }

	errOrNil := CV.Validate()
	if errOrNil != nil { return errOrNil }

	s.Repository.Save(CV)

	events := CV.PullEvents()
	CV.PublishEvent(s.Bus, events[0])
	return nil
}
