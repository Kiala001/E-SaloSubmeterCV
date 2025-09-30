package domain

import (
	"errors"

	"github.com/kindalus/godx/pkg/event"
)

type Candidate struct {
	Id       string
	Name     string
	Email    Email
	Password Password
	CVId     string
	events   []event.Event
}

func NewCandidate(id string, name string, email Email, password Password, cv_id string) (Candidate, error) {
	
	if name == "" {	return Candidate{}, errors.New("Name cannot be empty") }

	if cv_id == "" { return Candidate{}, errors.New("CVId cannot be empty") }

	Candidate := Candidate{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		CVId:     cv_id,
	}

	Candidate.AddEvent("CandidatoRegistadoCadastrado")
	return Candidate, nil
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
