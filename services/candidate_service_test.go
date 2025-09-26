package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestRegistarCadastrarCandidato(t *testing.T){
	t.Run("You must register a candidate successfully.", func (t *testing.T)  {
		eventBus := event.NewEventBus()
		candidato := domain.NewCandidato("Candidato001", "Silvano Varela", "CV001")
		
		candidatoRepository := adapters.NewInmemoryCandidatoRepository()
		candidatoService := NewCandidatoService(candidatoRepository, eventBus)
		candidatoService.RegistarCadastrarCandidato(candidato)

		length := candidatoRepository.Length()
		if length != 1 {
			t.Errorf("Espected %d, but got %d", 1, length)
		}
	})
	
	t.Run("You must publish the CandidatoRegistadoCadastrado event", func (t *testing.T)  {
		eventBus := event.NewEventBus()
		isPublished := false

		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CandidatoRegistadoCadastrado" {
				isPublished = true
			}
		})
		eventBus.Subscribe("CandidatoRegistadoCadastrado", eventHandler)
		
		candidato := domain.NewCandidato("Candidato002", "Rui Osvaldo", "CV002")

		candidatoRepository := adapters.NewInmemoryCandidatoRepository()
		candidatoService := NewCandidatoService(candidatoRepository, eventBus)
		candidatoService.RegistarCadastrarCandidato(candidato)

		if !isPublished {
			t.Error("I was hoping that the CandidatoRegistadoCadastrado event would be published.")
		}
	})

	t.Run("You must register candidate after submitting the CV.", func (t *testing.T)  {
		eventBus := event.NewEventBus()

		cvRepository := adapters.NewInmemoryCVRepository()
		cv := cvRepository.GetCVById("CV002")
		candidato := domain.NewCandidato("Candidato003", "Kiala Emanuel", cv.Id)

		cvService := NewCVSubmissionService(cvRepository, eventBus)

		candidatoRepository := adapters.NewInmemoryCandidatoRepository()
		candidatoService := NewCandidatoService(candidatoRepository, eventBus)
		
		var eventHandlerCVSubmetido = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				candidatoService.RegistarCadastrarCandidato(candidato)
			}
		})
		eventBus.Subscribe("CVSubmetido", eventHandlerCVSubmetido)
		
		cvService.SubmeterCV(cv)
		
		length := candidatoRepository.Length()
		if length != 1 {
			t.Errorf("Espected %d, but got %d", 1, length)
		}	
	})
	
	t.Run("You must not register candidate if the CV is not submitted.", func (t *testing.T)  {

		eventBus := event.NewEventBus()

		cvRepository := adapters.NewInmemoryCVRepository()
		cv := cvRepository.GetCVById("CV001")
		candidato := domain.NewCandidato("Candidato003", "Kiala Emanuel", cv.Id)

		cvService := NewCVSubmissionService(cvRepository, eventBus)

		candidatoRepository := adapters.NewInmemoryCandidatoRepository()
		candidatoService := NewCandidatoService(candidatoRepository, eventBus)
		
		var eventHandlerCVSubmetido = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				candidatoService.RegistarCadastrarCandidato(candidato)
			}
		})
		eventBus.Subscribe("CVSubmetido", eventHandlerCVSubmetido)
		
		cvService.SubmeterCV(cv)
		
		length := candidatoRepository.Length()
		if length != 0 {
			t.Errorf("Espected %d, but got %d", 0, length)
		}	
	})
}
