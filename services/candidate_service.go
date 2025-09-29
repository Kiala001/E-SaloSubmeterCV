package services

import (
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

func (s *candidateService) RegisterCandidate(CandidateDTO domain.CandidateDTO) error {
	email, err := domain.NewEmail(CandidateDTO.Email)
	if err != nil {	return err }

	password, err := domain.NewPassword(CandidateDTO.Password)
	if err != nil {	return err }

	Candidate, err := domain.NewCandidate(CandidateDTO.Id, CandidateDTO.Name, email, password, CandidateDTO.CVId)
	if err != nil {	return err	}

	s.repository.Save(Candidate)

	event := Candidate.PullEvents()

	Candidate.PublishEvents(s.eventBus, event[0])
	return nil
}
