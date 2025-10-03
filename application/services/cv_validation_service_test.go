package services

import (
	"esalo/adapters"
	"esalo/domain"
	"testing"
	
	"github.com/kindalus/godx/pkg/event"
)

func TestFeatureValidateCV(t *testing.T) {
	t.Run("Should publish the CVValidado event", func(t *testing.T) {
	
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
			t.Errorf("Expected CVValidado event to be published.")
		}
	})

	t.Run("Should change the CV status to Validado.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, EventBus)

		CVService.ValidateCV("CV001")

		updatedCV, _ := CVRepository.GetById("CV001")
		if updatedCV.Status() != domain.VALIDADO {
			t.Errorf("Expected %s, but got %s", domain.VALIDADO, updatedCV.Status())
		}
	})

	t.Run("Should not validate the CV if it is already validated.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, EventBus)

		err := CVService.ValidateCV("CV002")

		if err != domain.ErrCVAlreadyValidated {
			t.Errorf("Expected error %s, but got %s", domain.ErrCVAlreadyValidated, err)
		}
	})

	t.Run("Should return an error if the CV does not exist.", func(t *testing.T) {

		EventBus := event.NewEventBus()
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, EventBus)

		err := CVService.ValidateCV("CV999")

		if err != ErrCVNotFound {
			t.Errorf("Expected an error, but got nil")
		}
	})


	t.Run("Should return an error when trying to validate a CV that is already submitted.", func(t *testing.T) {
		EventBus := event.NewEventBus()
		
		CVRepository := adapters.NewInmemoryCVRepository()
		CVService := NewCVValidationService(CVRepository, EventBus)

		err := CVService.ValidateCV("CV003")
		
		if err != domain.ErrCannotValidateCV {
			t.Errorf("Expected %v, but got %v", domain.ErrCannotValidateCV, err)
		}
	})
}