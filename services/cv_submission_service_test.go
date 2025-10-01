package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestSubmissionCV(t *testing.T) {
	t.Run("You must publish the CVSubmetido event", func(t *testing.T) {
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
			t.Errorf("I was hoping that the CVSubmetido event would be published.")
		}
	})

	t.Run("You must Save th CV status to Available.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVSubmissionService(CVRepository, EventBus)
		CVService.SubmitCV("CV002")

		updatedCV, _ := CVRepository.GetById("CV002")

		if updatedCV.Status() != domain.Submetido {
			t.Errorf("Expected %s, but got %s", domain.Submetido, updatedCV.Status())
		}
	})

	t.Run("You must not submit the CV if it is not validated.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVSubmissionService(CVRepository, EventBus)
		ErrOrNil := CVService.SubmitCV("CV001")

		if ErrOrNil == nil {
			t.Errorf("Expected nil, but got %v", ErrOrNil)
		}
	})

	t.Run("Must not submit the CV if it is already submitted.", func(t *testing.T) {
		EventBus := event.NewEventBus()
		
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)
		
		ErrOrNil := CVService.SubmitCV("CV003")
		if ErrOrNil == nil {
			t.Errorf("Expected error, but got %v", nil)
		}
	})

}
