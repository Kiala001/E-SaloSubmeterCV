package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestRegisterCandidate(t *testing.T) {
	t.Run("You must register a candidate successfully.", func(t *testing.T) {
		eventBus := event.NewEventBus()

		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV001",
		}

		candidateRepository := adapters.NewInmemorycandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)
		candidateService.RegisterCandidate(Candidate)

		length := candidateRepository.Length()
		if length != 1 {
			t.Errorf("Expected %d, but got %d", 1, length)
		}
	})

	t.Run("You must not register a candidate with an empty cvid.", func(t *testing.T) {
		eventBus := event.NewEventBus()
		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "",
		}

		candidateRepository := adapters.NewInmemorycandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)
		ErrOrNil := candidateService.RegisterCandidate(Candidate)

		if ErrOrNil == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})

	t.Run("You must publish the CandidateRegistadoCadastrado event", func(t *testing.T) {
		eventBus := event.NewEventBus()
		isPublished := false

		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CandidatoRegistadoCadastrado" {
				isPublished = true
			}
		})
		eventBus.Subscribe("CandidatoRegistadoCadastrado", eventHandler)

		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV001",
		}

		candidateRepository := adapters.NewInmemorycandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)
		candidateService.RegisterCandidate(Candidate)

		if !isPublished {
			t.Error("I was hoping that the CandidateRegistadoCadastrado event would be published.")
		}
	})

	t.Run("You must register a candidate after validating the CV.", func(t *testing.T) {
		eventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVValidationService(CVRepository, eventBus)

		candidateRepository := adapters.NewInmemorycandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
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

	t.Run("You must register candidate after submitting the CV.", func(t *testing.T) {
		eventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()

		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV002",
		}

		CVService := NewCVSubmissionService(CVRepository, eventBus)

		candidateRepository := adapters.NewInmemorycandidateRepository()
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

	t.Run("You must not register candidate if the CV is not submitted.", func(t *testing.T) {

		eventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()

		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV001",
		}

		CVService := NewCVSubmissionService(CVRepository, eventBus)

		candidateRepository := adapters.NewInmemorycandidateRepository()
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

	t.Run("You must register as a candidate after submitting your CV", func(t *testing.T) {
		eventBus := event.NewEventBus()
		isPublished := false

		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVSubmissionService(CVRepository, eventBus)

		Candidate := domain.CandidateDTO{
			Id:       "Candidate003",
			Name:     "Kiala Emanuel",
			Email:    "kiala@gmail.com",
			Password: "Kiala001",
			CVId:     "CV002",
		}

		CandidateRepository := adapters.NewInmemorycandidateRepository()
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
			t.Errorf("I was hoping that the CandidatoRegistadoCadastrado event would be published.")
		}
	})
}
