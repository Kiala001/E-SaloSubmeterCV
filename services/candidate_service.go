package services

import (
	"esalo/domain"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)

type candidateService struct {
	repository ports.IcandidateRepository
	eventBus   event.Bus
}

func NewcandidateService(repo ports.IcandidateRepository, eventBus event.Bus) *candidateService {
	return &candidateService{repository: repo, eventBus: eventBus}
}

func (s *candidateService) RegisterCandidate(CandidateDTO domain.CandidateDTO) error {
	Candidate, err := domain.NewCandidate(CandidateDTO.Id, CandidateDTO.Name, CandidateDTO.Email, CandidateDTO.Password, CandidateDTO.CVId)
	if err != nil {
		return err
	}

	s.repository.Save(Candidate)

	event := Candidate.PullEvents()

	Candidate.PublishEvents(s.eventBus, event[0])
	return nil
}
