package tests

import (
	"esalo/adapters"
	"esalo/services"
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestSubmeterCV(t *testing.T) {
	t.Run("You must publish the CVSubmetido event", func (t *testing.T)  {
		EventBus := event.NewEventBus()
		isPublished := false
		
		var eventHandler = event.HandlerFunc(func (e event.Event){
			if e.Name() == "CVSubmetido" {
				isPublished = true
			}
		})
		EventBus.Subscribe("CVSubmetido", eventHandler)


		CVRepository := adapters.NewInmemoryCVRepository()
		cv := CVRepository.GetCVById("CV002")

		CVService := services.NewCVSubmissionService(CVRepository, EventBus)
		CVService.SubmeterCV(cv)

		if !isPublished {
			t.Errorf("I was hoping that the CVSubmetido event would be published.")
		}
	})
	
	t.Run("You must update th CV status to Available.", func (t *testing.T)  {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()
		cv := CVRepository.GetCVById("CV002")

		CVService := services.NewCVSubmissionService(CVRepository, EventBus)
		CVService.SubmeterCV(cv)

		CvActualizado := CVRepository.GetCVById(cv.Id)

		if CvActualizado.Estado != "Disponível" {
			t.Errorf("Espected %s, but got %s", "Disponível", CvActualizado.Estado)
		}
	})

	t.Run("You must not submit the CV if it is not validated.", func (t *testing.T)  {
		EventBus := event.NewEventBus()

		CVRepository := adapters.NewInmemoryCVRepository()
		cv := CVRepository.GetCVById("CV001")

		CVService := services.NewCVSubmissionService(CVRepository, EventBus)
		ErrOrNil := CVService.SubmeterCV(cv)

		if ErrOrNil != nil {
			t.Errorf("Espected nil, but got %v", ErrOrNil)
		}
	})
}