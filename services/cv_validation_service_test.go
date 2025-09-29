package services

import (
	"esalo/adapters"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestValidatedCV(t *testing.T) {

	t.Run("You must publish the CVValidado event", func(t *testing.T) {
		EventBus := event.NewEventBus()
		isPublished := false

		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVValidado" {
				isPublished = true
			}
		})
		EventBus.Subscribe("CVValidado", eventHandler)

		CVRepository := adapters.NewInmemoryCVRepository()
		cv := CVRepository.GetCVById("CV001")

		CVService := NewCVValidationService(CVRepository, EventBus)
		CVService.ValidateCV(cv)

		if !isPublished {
			t.Errorf("I was hoping that the CVValidado event would be published.")
		}
	})

	t.Run("You must update the CV status to validated.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()
		cv := CVRepository.GetCVById("CV001")

		CVService := NewCVValidationService(CVRepository, EventBus)
		CVService.ValidateCV(cv)

		UpdatedCV := CVRepository.GetCVById(cv.Id)

		if UpdatedCV.Status != "Validado" {
			t.Errorf("Expected %s, but got %s", "Validado", UpdatedCV.Status)
		}
	})

}
