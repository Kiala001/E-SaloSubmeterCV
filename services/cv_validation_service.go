package services

import (
	"esalo/domain"
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

func (s *CVValidationService) ValidateCV(cv domain.CV) error {
	cv.Validate()

	s.Repository.Update(cv)

	events := cv.PullEvents()
	cv.PublishEvent(s.Bus, events[0])
	return nil
}
