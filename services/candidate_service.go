package services

import (
	"errors"
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

func (s *candidateService) RegisterCandidate(CandidateDTO application.CandidateDTO) error {

	name, errOrNil := domain.NewName(CandidateDTO.Name)
	if errOrNil != nil { 
		return errOrNil 
	}

	email, errOrNil := domain.NewEmail(CandidateDTO.Email)
	if errOrNil != nil {	
		return errOrNil 
	}

	password, errOrNil := domain.NewPassword(CandidateDTO.Password)
	if errOrNil != nil {	
		return errOrNil 
	}

	if _, exists := s.repository.FindByEmail(email); exists { 
		return errors.New("candidate with this email already exists") 
	}

	candidate, errOrNil := domain.NewCandidate(name, email, password, CandidateDTO.CVId)
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