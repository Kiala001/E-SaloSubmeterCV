package ports

import "esalo/domain"


type ICandidatoRepository interface {
	Length() int
	Save(candidato *domain.Candidato) error
}