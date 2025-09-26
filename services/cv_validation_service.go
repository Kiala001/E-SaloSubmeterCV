package services

import (
	"esalo/domain"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type CVValidationService struct {
	Bus        event.Bus
	Repository ports.ICVRepository
}

func NewCVValidationService(repo ports.ICVRepository, bus event.Bus) *CVValidationService {
	return &CVValidationService{Repository: repo, Bus: bus}
}

func (s *CVValidationService) ValidarCV(cv domain.CV) error {
	cv.Validar()

	s.Repository.Update(cv)

	e := event.New("CVValidado", event.WithPayload(cv))	
	cv.PublishEvent(s.Bus, e)
	return nil
}
