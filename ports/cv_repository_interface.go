package ports

import "esalo/domain"

type ICVRepository interface {
	Update(cv domain.CV) error
	GetCVById(id string) domain.CV
}
