package domain

import (
	"errors"

	"github.com/kindalus/godx/pkg/event"
)

type CV struct {
	Id     string
	Status string
	events []event.Event
}

func (c *CV) Validate() error {
	if c.Status == "Validado" {
		return errors.New("CV already validated")
	}

	c.Status = "Validado"

	c.AddEvent("CVValidado")
	return nil
}

func (c *CV) Submit() error {
	if c.Status != "Validado" {
		return errors.New("CV must be validated before submission")
	}

	if c.Status == "Submetido" {
		return errors.New("CV already submitted")
	}

	c.Status = "Submetido"

	c.AddEvent("CVSubmetido")
	return nil
}

func (c *CV) PublishEvent(bus event.Bus, e event.Event) {
	bus.Publish(e)
}

func (c *CV) AddEvent(eventName string) {
	c.events = append(c.events, event.New(eventName, event.WithPayload(c)))
}

func (c *CV) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}
	return events
}
