package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestValidatedCV(t *testing.T) {

	t.Run("Must publish the CVValidado event", func(t *testing.T) {
		EventBus := event.NewEventBus()
		isPublished := false

		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVValidado" {
				isPublished = true
			}
		})
		EventBus.Subscribe("CVValidado", eventHandler)

		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVValidationService(CVRepository, EventBus)
		CVService.ValidateCV("CV001")

		if !isPublished {
			t.Errorf("I was hoping that the CVValidado event would be published.")
		}
	})

	t.Run("Must Save the CV status to validated.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()

		CVService := NewCVValidationService(CVRepository, EventBus)
		CVService.ValidateCV("CV001")

		updatedCV, _ := CVRepository.GetById("CV001")

		if updatedCV.Status() != domain.Validado {
			t.Errorf("Expected %s, but got %s", domain.Validado, updatedCV.Status())
		}
	})

	t.Run("Must not validate the CV if it is already validated.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, EventBus)

		CVService.ValidateCV("CV002")

		updatedCV, _ := CVRepository.GetById("CV002")
		if updatedCV.Status() != domain.Validado {
			t.Errorf("Expected %s, but got %s", domain.Validado, updatedCV.Status())
		}
	})

	t.Run("Must return an error if the CV does not exist.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, EventBus)

		err := CVService.ValidateCV("CV999")

		if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
		
	})
}
