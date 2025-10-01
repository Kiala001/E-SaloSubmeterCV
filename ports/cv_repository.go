package ports

import "esalo/domain"

type CVRepository interface {
	Save(cv domain.CV) error
	GetById(id string) (domain.CV, bool)
}
