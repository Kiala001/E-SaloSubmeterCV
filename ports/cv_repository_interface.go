package ports

import "esalo/domain"

type CVRepository interface {
	Update(cv domain.CV) error
	GetCVById(id string) domain.CV
}
