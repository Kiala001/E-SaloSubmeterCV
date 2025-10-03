package services

import (
	"esalo/application"
	"esalo/domain"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type candidateService struct {
	repository ports.CandidateRepository
	eventBus   event.Bus
}

func NewcandidateService(repo ports.CandidateRepository, eventBus event.Bus) *candidateService {
	return &candidateService{repository: repo, eventBus: eventBus}
}

func (s *candidateService) RegisterCandidate(CandidateData application.CandidateData) error {
	cvId, errOrNil := domain.NewCVId(CandidateData.CVId)
	if errOrNil != nil {
		return errOrNil
	}

	name, errOrNil := domain.NewName(CandidateData.Name)
	if errOrNil != nil {
		return errOrNil
	}
	
	password, errOrNil := domain.NewPassword(CandidateData.Password)
	if errOrNil != nil {
		return errOrNil
	}
	
	email, errOrNil := domain.NewEmail(CandidateData.Email)
	if errOrNil != nil {
		return errOrNil
	}

	if _, exists := s.repository.FindByEmail(email); exists {
		return ErrCandidateAlreadyExists
	}

	candidate, errOrNil := domain.NewCandidate(name, email, password, cvId)
	if errOrNil != nil {
		return errOrNil
	}

	s.repository.Save(candidate)

	events := candidate.PullEvents()
	for _, event := range events {
		s.eventBus.Publish(event)
	}

	return nil
}
