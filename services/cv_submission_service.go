package services


import (
	"esalo/domain"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type CVSubmissionService struct {
	Bus        event.Bus
	Repository ports.ICVRepository
}

func NewCVSubmissionService(repo ports.ICVRepository, bus event.Bus) *CVSubmissionService {
	return &CVSubmissionService{Repository: repo, Bus: bus}
}

func (s *CVSubmissionService) SubmeterCV(cv domain.CV) error {
	if cv.Estado != "Válido" { return nil }

	cv.Submeter()

	s.Repository.Update(cv)

	e := event.New("CVSubmetido", event.WithPayload(cv))
	cv.PublishEvent(s.Bus, e)
	return nil
}