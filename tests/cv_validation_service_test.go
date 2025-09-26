package tests

import (
	"esalo/adapters"
	"esalo/services"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestValidarCV(t *testing.T) {

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

		CVService := services.NewCVValidationService(CVRepository, EventBus)
		CVService.ValidarCV(cv)

		if !isPublished {
			t.Errorf("I was hoping that the CVValidado event would be published.")
		}
	})
	
	t.Run("You must update the CV status to valid.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()
		cv := CVRepository.GetCVById("CV001")

		CVService := services.NewCVValidationService(CVRepository, EventBus)
		CVService.ValidarCV(cv)

		CvActualizado := CVRepository.GetCVById(cv.Id)

		if CvActualizado.Estado != "Válido" {
			t.Errorf("Espected %s, but got %s", "Valido", CvActualizado.Estado)
		}
	})
}
