package adapters

import "esalo/domain"

type InmemoryCVRepository struct {
	CVs map[string]domain.CV
}

func NewInmemoryCVRepository() *InmemoryCVRepository {
	return &InmemoryCVRepository{CVs: map[string]domain.CV{
		"CV001": {Id: "CV001", Estado: "Ativo"},
		"CV002": {Id: "CV002", Estado: "Válido"},
	}}
}

func (r *InmemoryCVRepository) GetCVById(id string) domain.CV {
	return r.CVs[id]
}

func (r *InmemoryCVRepository) Update(cv domain.CV) error {
	r.CVs[cv.Id] = cv
	return nil
}