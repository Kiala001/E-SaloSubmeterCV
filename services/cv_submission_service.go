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

func (s *CVSubmissionService) SubmeterCV(CV domain.CV) error {

	error := CV.Submeter()
	if error != nil {
		return error
	}

	s.Repository.Update(CV)
	
	events := CV.PullEvents()
	CV.PublishEvent(s.Bus, events[0])
	
	return nil
}