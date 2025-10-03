package domain

import (
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
	if c.status == VALIDADO {
		return ErrCVAlreadyValidated
	}

	if c.status == SUBMETIDO {
		return ErrCannotValidateCV
	}

	c.status = VALIDADO
	c.AddEvent("CVValidado")
	return nil
}

func (c *CV) Submit() error {
	if c.status == SUBMETIDO {
		return ErrCVAlreadySubmitted
	}

	if c.status != VALIDADO {
		return ErrCannotSubmitCV
	}

	c.status = SUBMETIDO
	c.AddEvent("CVSubmetido")
	return nil
}

func (c CV) Status() CVStatus {
	return c.status
}

func (c CV) Id() string {
	return c.id
}

func (c *CV) AddEvent(eventName string) {
	c.events = append(c.events, event.New(eventName, event.WithPayload(c)))
}

func (c *CV) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}
	return events
}
