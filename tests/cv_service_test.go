package tests

import (
	"testing"

	"github.com/kindalus/godx/pkg/event"
)

func TestValidarCV(t *testing.T) {

	t.Run("You must publish the CVValidado event", func(t *testing.T) {
		EventBus := event.NewEventBus()
		isPublished := false

		cv := CV{
			Id:     "CV001",
			Estado: "Ativo",
		}

		var eventHandler = event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVValidado" {
				isPublished = true
			}
		})

		EventBus.Subscribe("CVValidado", eventHandler)

		CVRepository := NewCVRepository()

		CVService := NewCVService(CVRepository, EventBus)
		CVService.ValidarCV(cv)

		if !isPublished {
			t.Errorf("I was hoping that the CVValidado event would be published.")
		}
	})
	
	t.Run("You must update the CV status to valid.", func(t *testing.T) {
		EventBus := event.NewEventBus()

		cv := CV{
			Id:     "CV002",
			Estado: "Ativo",
		}
		CVRepository := NewCVRepository()

		CVService := NewCVService(CVRepository, EventBus)
		CVService.ValidarCV(cv)

		CvActualizado := CVRepository.FindById(cv.Id)

		if CvActualizado.Estado != "Válido" {
			t.Errorf("Espected %s, but got %s", "Valido", CvActualizado.Estado)
		}
	})

}

func TestSubmeterCV(t *testing.T) {
	t.Run("You must publish the CVSubmetido event", func (t *testing.T)  {
		EventBus := event.NewEventBus()
		isPublished := false
		cv := CV{
			Id:     "CV003",
			Estado: "Válido",
		}
		
		var eventHandler = event.HandlerFunc(func (e event.Event){
			if e.Name() == "CVSubmetido" {
				isPublished = true
			}
		})
		EventBus.Subscribe("CVSubmetido", eventHandler)


		CVRepository := NewCVRepository()
		CVService := NewCVService(CVRepository, EventBus)
		CVService.SubmeterCV(cv)

		if !isPublished {
			t.Errorf("I was hoping that the CVSubmetido event would be published.")
		}
	})
	
	t.Run("You must update th CV status to Available.", func (t *testing.T)  {
		EventBus := event.NewEventBus()
		cv := CV {
			Id:     "CV004",
			Estado: "Válido",
		}

		CVRepository := NewCVRepository()
		CVService := NewCVService(CVRepository, EventBus)
		CVService.SubmeterCV(cv)

		CvActualizado := CVRepository.FindById(cv.Id)

		if CvActualizado.Estado != "Disponível" {
			t.Errorf("Espected %s, but got %s", "Disponível", CvActualizado.Estado)
		}
	})

	t.Run("You must not submit the CV if it is not validated.", func (t *testing.T)  {
		EventBus := event.NewEventBus()
		cv := CV {
			Id:    "CV005",
			Estado: "Ativo",
		}

		CVRepository := NewCVRepository()
		CVService := NewCVService(CVRepository, EventBus)
		ErrOrNil := CVService.SubmeterCV(cv)

		if ErrOrNil != nil {
			t.Errorf("Espected nil, but got %v", ErrOrNil)
		}
	})
}

type CV struct {
	Id     string
	Estado string
}

func (c *CV) Validar() error {
	c.Estado = "Válido"
	return nil
}

func (c *CV) Submeter() error {
	c.Estado = "Disponível"
	return nil
}

func (c *CV) PublishEvent(bus event.Bus, e event.Event) {
	bus.Publish(e)
}

type CVService struct {
	Bus        event.Bus
	Repository ICVRepository
}

func NewCVService(repo ICVRepository, bus event.Bus) *CVService {
	return &CVService{Repository: repo, Bus: bus}
}

func (s *CVService) ValidarCV(cv CV) error {
	cv.Validar()

	s.Repository.Update(cv)

	e := event.New("CVValidado", event.WithPayload(cv))	
	cv.PublishEvent(s.Bus, e)
	return nil
}

func (s *CVService) SubmeterCV(cv CV) error {
	if cv.Estado != "Válido" { return nil }

	cv.Submeter()

	s.Repository.Update(cv)

	e := event.New("CVSubmetido", event.WithPayload(cv))
	cv.PublishEvent(s.Bus, e)
	return nil
}

type ICVRepository interface {
	Update(cv CV) error
	FindById(id string) CV
}

type CVRepository struct {
	CVs map[string]CV
}

func NewCVRepository() *CVRepository {
	return &CVRepository{CVs: make(map[string]CV)}
}

func (r *CVRepository) FindById(id string) CV {
	return r.CVs[id]
}

func (r *CVRepository) Update(cv CV) error {
	r.CVs[cv.Id] = cv
	return nil
}