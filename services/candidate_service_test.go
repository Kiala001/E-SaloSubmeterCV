package services

import (
	"esalo/adapters"
	"esalo/application"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestFeatureRegisterCandidate(t *testing.T) {
	t.Run("Should register a candidate successfully.", func(t *testing.T) {
		Candidate := application.CandidateDTO{
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV001",
		}

		eventBus := event.NewEventBus()
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		candidateService.RegisterCandidate(Candidate)

		length := candidateRepository.Length()
		if length != 1 {
			t.Errorf("Expected %d, but got %d", 1, length)
		}
	})

	t.Run("Should publish the CandidateRegistadoCadastrado event", func(t *testing.T) {
		Candidate := application.CandidateDTO{
			Name:     "Rui Manuel",
			Email:    "rui@gmail.com",
			Password: "Rui001",
			CVId:     "CV001",
		}

		eventBus := event.NewEventBus()
		isPublished := false
		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CandidatoRegistadoCadastrado" {
				isPublished = true
			}
		})
		eventBus.Subscribe("CandidatoRegistadoCadastrado", eventHandler)
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		candidateService.RegisterCandidate(Candidate)

		if !isPublished {
			t.Error("Expected CandidateRegistadoCadastrado event to be published.")
		}
	})

	t.Run("Should not register the second candidate with the same email address as the first candidate.", func(t *testing.T) {
		firstCandidate := application.CandidateDTO{
			Name:     "Mário Varela",
			Email:    "esalo@gmail.com",
			Password: "Varela001",
			CVId:     "CV001",
		}

		secondCandidate := application.CandidateDTO{
			Name:     "Silvano Paulino",
			Email:    "esalo@gmail.com",
			Password: "Silvano001",
			CVId:     "CV002",
		}

		eventBus := event.NewEventBus()
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		candidateService.RegisterCandidate(firstCandidate)
		candidateService.RegisterCandidate(secondCandidate)

		length := candidateRepository.Length()
		if length != 1 {
			t.Errorf("Expected %d, but got %d", 1, length)
		}

	})

	t.Run("Should not register a candidate with an empty cvid.", func(t *testing.T) {

		eventBus := event.NewEventBus()
		Candidate := application.CandidateDTO{
			Name:     "Osvaldo de Sousa",
			Email:    "osvaldo@gmail.com",
			Password: "osvaldo001",
			CVId:     "",
		}
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		ErrOrNil := candidateService.RegisterCandidate(Candidate)

		if ErrOrNil == nil {
			t.Error("Expecting an error, but I got nil.")
		}
	})

	t.Run("Should register a candidate after validating the CV.", func(t *testing.T) {

		eventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, eventBus)
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		Candidate := application.CandidateDTO{
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV001",
		}
		var eventHandlerCVSubmetido = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVValidado" {
				candidateService.RegisterCandidate(Candidate)
			}
		})
		eventBus.Subscribe("CVValidado", eventHandlerCVSubmetido)

		CVService.ValidateCV("CV001")

		length := candidateRepository.Length()
		if length != 1 {
			t.Errorf("Expected %d, but got %d", 1, length)
		}
	})

	t.Run("Should register candidate after submitting the CV.", func(t *testing.T) {

		Candidate := application.CandidateDTO{
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV002",
		}
		eventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVSubmissionService(CVRepository, eventBus)
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		var eventHandlerCVSubmetido = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				candidateService.RegisterCandidate(Candidate)
			}
		})
		eventBus.Subscribe("CVSubmetido", eventHandlerCVSubmetido)

		CVService.SubmitCV("CV002")

		length := candidateRepository.Length()
		if length != 1 {
			t.Errorf("Expected %d, but got %d", 1, length)
		}
	})

	t.Run("Should not register candidate if the CV is not submitted.", func(t *testing.T) {

		Candidate := application.CandidateDTO{
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV001",
		}

		eventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, eventBus)
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		var eventHandlerCVSubmetido = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				candidateService.RegisterCandidate(Candidate)
			}
		})
		eventBus.Subscribe("CVSubmetido", eventHandlerCVSubmetido)

		CVService.SubmitCV("CV001")

		length := candidateRepository.Length()
		if length != 0 {
			t.Errorf("Expected %d, but got %d", 0, length)
		}
	})

	t.Run("Should register as a candidate after submitting your CV", func(t *testing.T) {

		Candidate := application.CandidateDTO{
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV002",
		}

		eventBus := event.NewEventBus()
		isPublished := false
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, eventBus)

		CandidateRepository := adapters.NewInmemoryCandidateRepository()
		CandidateService := NewcandidateService(CandidateRepository, eventBus)

		var eventHandlerCVSubmetido = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				CandidateService.RegisterCandidate(Candidate)
			}
		})

		var eventHandlerCandidateRegistadoCadastrado = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CandidatoRegistadoCadastrado" {
				isPublished = true
			}
		})
		eventBus.Subscribe("CVSubmetido", eventHandlerCVSubmetido)
		eventBus.Subscribe("CandidatoRegistadoCadastrado", eventHandlerCandidateRegistadoCadastrado)

		CVService.SubmitCV("CV002")

		if !isPublished {
			t.Errorf("Expected CandidatoRegistadoCadastrado event to be published.")
		}
	})

}
