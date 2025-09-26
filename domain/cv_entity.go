package domain

import (
	"errors"

	"github.com/kindalus/godx/pkg/event"
)

type CV struct {
	Id     string
	Estado string
	events []event.Event
}

func (c *CV) Validar() error {
	c.Estado = "Validado"

	c.events = append(c.events, event.New("CVValidado", event.WithPayload(c)))
	return nil
}

func (c *CV) Submeter() error {
	if c.Estado != "Validado" {
		return errors.New("CV must be validated before submission")
	}

	c.Estado = "Submetido"

	c.events = append(c.events, event.New("CVSubmetido", event.WithPayload(c)))
	return nil
}

func (c *CV) PublishEvent(bus event.Bus, e event.Event) {
	bus.Publish(e)
}

func (c *CV) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}
	return events
}