package domain

import (
	"errors"

	"github.com/kindalus/godx/pkg/event"
)

type Candidate struct {
	id       ID
	name     Name
	email    Email
	password Password
	cvId     string
	events   []event.Event
}

func NewCandidate(name Name, email Email, password Password, cvId string) (Candidate, error) {

	if cvId == "" {	return Candidate{}, errors.New("CVId cannot be empty") }

	id, err := NewID()
	if err != nil {	return Candidate{}, err }

	Candidate := Candidate{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		cvId:     cvId,
	}

	Candidate.AddEvent("CandidatoRegistadoCadastrado")
	return Candidate, nil
}

func (c Candidate) ID() ID {
	return c.id
}

func (c Candidate) Email() Email {
	return c.email
}

func (c *Candidate) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}
	return events
}

func (c *Candidate) AddEvent(eventName string) {
	c.events = append(c.events, event.New(eventName, event.WithPayload(c)))
}

func (c *Candidate) PublishEvents(bus event.Bus, e event.Event) {
	bus.Publish(e)
}