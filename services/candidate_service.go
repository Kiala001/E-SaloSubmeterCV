package services

import (
	"esalo/domain"
	"esalo/ports"

	"github.com/kindalus/godx/pkg/event"
)


type CandidatoService struct {
	repository ports.ICandidatoRepository
	eventBus event.Bus
}

func NewCandidatoService(repo ports.ICandidatoRepository, eventBus event.Bus) *CandidatoService {
	return &CandidatoService{repository: repo, eventBus: eventBus}
}

func (s *CandidatoService) RegistarCadastrarCandidato(candidato *domain.Candidato) error {
	s.repository.Save(candidato)	

	event := candidato.PullEvents()
	
	candidato.PublishEvents(s.eventBus, event[0])
	
	return nil
}