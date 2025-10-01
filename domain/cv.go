package domain

import (
	"errors"

	"github.com/kindalus/godx/pkg/event"
)

type CV struct {
	id     string
	status CVStatus
	events []event.Event
}

func NewCV(id string, status CVStatus) CV {
	return CV{
		id:     id,
		status: status,
	}
}

func (c *CV) Validate() error {
	if c.status == Validado {
		return errors.New("CV already validated")
	}

	if c.status == Submetido {
		return errors.New("CV already submitted and cannot be validated")		
	}

	c.status = Validado

	c.AddEvent("CVValidado")
	return nil
}

func (c *CV) Submit() error {
	if c.status != Validado {
		return errors.New("CV must be validated before submission")
	}

	if c.status == Submetido {
		return errors.New("CV already submitted")
	}

	c.status = Submetido

	c.AddEvent("CVSubmetido")
	return nil
}

func (c CV) Status() CVStatus {
	return c.status
}

func (c CV) Id() string {
	return c.id
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
