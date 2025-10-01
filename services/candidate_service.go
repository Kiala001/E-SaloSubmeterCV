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

func (s *candidateService) RegisterCandidate(CandidateDTO application.CandidateDTO) error {
	name, err := domain.NewName(CandidateDTO.Name)
	if err != nil {	return err }

	email, err := domain.NewEmail(CandidateDTO.Email)
	if err != nil {	return err }

	password, err := domain.NewPassword(CandidateDTO.Password)
	if err != nil {	return err }

	_, exists := s.repository.FindByEmail(email)
	if exists { return err }

	Candidate, err := domain.NewCandidate(name, email, password, CandidateDTO.CVId)
	if err != nil {	return err	}

	s.repository.Save(Candidate)
	
	event := Candidate.PullEvents()

	Candidate.PublishEvents(s.eventBus, event[0])
	return nil
}