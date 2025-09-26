package domain

import "github.com/kindalus/godx/pkg/event"

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