package adapters

import "esalo/domain"

type InmemoryCVRepository struct {
	CVs map[string]domain.CV
}

func NewInmemoryCVRepository() *InmemoryCVRepository {
	return &InmemoryCVRepository{CVs: map[string]domain.CV{
		"CV001": domain.NewCV("CV001", domain.Criado),
		"CV002": domain.NewCV("CV002", domain.Validado),
		"CV003": domain.NewCV("CV003", domain.Submetido),
	}}
}

func (r *InmemoryCVRepository) GetById(id string) (domain.CV, bool) {
	cv, ok := r.CVs[id]
	return cv, ok
}

func (r *InmemoryCVRepository) Save(cv domain.CV) error {
	r.CVs[cv.Id()] = cv
	return nil
}
