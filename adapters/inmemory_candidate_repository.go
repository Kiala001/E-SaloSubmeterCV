package adapters

import "esalo/domain"


type InmemoryCandidatoRepository struct {
	Candidatos map[string]*domain.Candidato
}

func NewInmemoryCandidatoRepository() *InmemoryCandidatoRepository {
	return &InmemoryCandidatoRepository{
		Candidatos: make(map[string]*domain.Candidato),
	}
}

func (r *InmemoryCandidatoRepository) Length() int {
	return len(r.Candidatos)
}

func (r *InmemoryCandidatoRepository) Save(candidato *domain.Candidato) error {
	r.Candidatos[candidato.Id] = candidato
	return nil
}