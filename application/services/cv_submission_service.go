package services

import (
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type CVSubmissionService struct {
	bus        event.Bus
	repository ports.CVRepository
}

func NewCVSubmissionService(repo ports.CVRepository, bus event.Bus) *CVSubmissionService {
	return &CVSubmissionService{repository: repo, bus: bus}
}

func (s *CVSubmissionService) SubmitCV(CvId string) error {
	CV, exists := s.repository.GetById(CvId)
	if !exists { 
		return ErrCVNotFound 
	}

	if errOrNil := CV.Submit(); errOrNil != nil { 
		return errOrNil 
	}

	s.repository.Save(CV)

	events := CV.PullEvents()

	for _, event := range events {
		s.bus.Publish(event)
	}

	return nil
}