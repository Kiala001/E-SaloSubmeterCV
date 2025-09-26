package domain

import "github.com/kindalus/godx/pkg/event"


type Candidato struct {
	Id string
	Nome string
	CVId string
	events []event.Event
}

func NewCandidato(id string, nome string, cv_id string) *Candidato {
	candidato := &Candidato{
		Id: id,
		Nome: nome,
		CVId: cv_id,
	}

	candidato.events = append(candidato.events, event.New("CandidatoRegistadoCadastrado", event.WithPayload(candidato)))
	return candidato
}

func (c *Candidato) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}
	return events
}

func (c *Candidato) PublishEvents(bus event.Bus, e event.Event) {
	bus.Publish(e)
}