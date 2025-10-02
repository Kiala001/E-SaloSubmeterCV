package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestFeatureSubmissionCV(t *testing.T) {

	t.Run("Should Save the CV status to validated.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		CVService.SubmitCV("CV002")

		updatedCV, _ := CVRepository.GetById("CV002")
		if updatedCV.Status() != domain.Submetido {
			t.Errorf("Expected %s, but got %s", domain.Submetido, updatedCV.Status())
		}
	})

	t.Run("Should not submit the CV if it is not validated.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		ErrOrNil := CVService.SubmitCV("CV001")

		if ErrOrNil == nil {
			t.Errorf("Expected nil, but got %v", ErrOrNil)
		}
	})

	t.Run("Should not submit the CV if it is already submitted.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		ErrOrNil := CVService.SubmitCV("CV003")

		if ErrOrNil == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})

	t.Run("Should return an error if the CV does not exist.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		err := CVService.SubmitCV("CV111")

		if err == nil {
			t.Errorf("Expected 'CV not found' error, but got nil")
		}
	})

	t.Run("Should publish the CVSubmetido event", func(t *testing.T) {

		EventBus := event.NewEventBus()
		isPublished := false
		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				isPublished = true
			}
		})
		EventBus.Subscribe("CVSubmetido", eventHandler)
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		CVService.SubmitCV("CV002")

		if !isPublished {
			t.Errorf("Expected CVSubmetido event to be published.")
		}
	})

	t.Run("Should not publish the CVSubmetido event if the CV is not submited.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		isPublished := false
		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVSubmetido" {
				isPublished = true
			}
		})
		EventBus.Subscribe("CVSubmetido", eventHandler)
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		CVService.SubmitCV("CV001")

		if isPublished {
			t.Errorf("Expected CVSubmetido event would not be published.")
		}
	})
}
