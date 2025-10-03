package services

import (
	"esalo/adapters"
	"esalo/application"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestFeatureRegisterCandidate(t *testing.T) {
	t.Run("Should register a candidate successfully.", func(t *testing.T) {
		Candidate := application.CandidateData{
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

	t.Run("Should publish the CandidatoRegistadoCadastrado event", func(t *testing.T) {
		Candidate := application.CandidateData{
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
			t.Error("Expected CandidatoRegistadoCadastrado event to be published.")
		}
	})

	t.Run("Should return an error when trying to register a candidate with an already existing email.", func(t *testing.T) {
		firstCandidate := application.CandidateData{
			Name:     "Mário Varela",
			Email:    "esalo@gmail.com",
			Password: "Varela001",
			CVId:     "CV001",
		}

		secondCandidate := application.CandidateData{
			Name:     "Silvano Paulino",
			Email:    "esalo@gmail.com",
			Password: "Silvano001",
			CVId:     "CV002",
		}

		eventBus := event.NewEventBus()
		candidateRepository := adapters.NewInmemoryCandidateRepository()
		candidateService := NewcandidateService(candidateRepository, eventBus)

		candidateService.RegisterCandidate(firstCandidate)
		err := candidateService.RegisterCandidate(secondCandidate)

		if err != ErrCandidateAlreadyExists {
			t.Errorf("Expected error %v, but got %v", ErrCandidateAlreadyExists, err)
		}

	})

	t.Run("Should not register a candidate with an empty cvid.", func(t *testing.T) {

		eventBus := event.NewEventBus()
		Candidate := application.CandidateData{
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
}
