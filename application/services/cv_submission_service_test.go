package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestFeatureSubmissionCV(t *testing.T) {

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

	t.Run("Should change the CV status to 'Submetido'.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		CVService.SubmitCV("CV002")

		updatedCV, _ := CVRepository.GetById("CV002")
		if updatedCV.Status() != domain.SUBMETIDO {
			t.Errorf("Expected %s, but got %s", domain.SUBMETIDO, updatedCV.Status())
		}
	})

	t.Run("Should not submit the CV if it is not validated.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		ErrOrNil := CVService.SubmitCV("CV001")

		if ErrOrNil != domain.ErrCannotSubmitCV {
			t.Errorf("Expected error %v, but got %v", domain.ErrCannotSubmitCV, ErrOrNil)
		}
	})

	t.Run("Should not submit the CV if it is already submitted.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		ErrOrNil := CVService.SubmitCV("CV003")

		if ErrOrNil != domain.ErrCVAlreadySubmitted  {
			t.Errorf("Expected error %v, but got %v", domain.ErrCVAlreadySubmitted, ErrOrNil)
		}
	})

	t.Run("Should return an error if the CV does not exist.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVSubmissionService(CVRepository, EventBus)

		err := CVService.SubmitCV("CV111")

		if err != ErrCVNotFound {
			t.Errorf("Expected an error %v, but got %v", ErrCVNotFound, err)
		}
	})
}
