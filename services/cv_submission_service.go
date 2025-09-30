package services

import (
	"errors"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type CVSubmissionService struct {
	Bus        event.Bus
	Repository ports.CVRepository
}

func NewCVSubmissionService(repo ports.CVRepository, bus event.Bus) *CVSubmissionService {
	return &CVSubmissionService{Repository: repo, Bus: bus}
}

func (s *CVSubmissionService) SubmitCV(CvId string) error {
	CV, exists := s.Repository.GetById(CvId)
	if !exists { return errors.New("CV not found") }

	errOrNil := CV.Submit()
	if errOrNil != nil {return errOrNil	}

	s.Repository.Save(CV)

	events := CV.PullEvents()
	CV.PublishEvent(s.Bus, events[0])

	return nil
}
