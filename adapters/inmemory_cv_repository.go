package adapters

import "esalo/domain"

type InmemoryCVRepository struct {
	cvs map[string]domain.CV
}

func NewInmemoryCVRepository() *InmemoryCVRepository {
	return &InmemoryCVRepository{cvs: map[string]domain.CV{
		"CV001": domain.NewCV("CV001", domain.Criado),
		"CV002": domain.NewCV("CV002", domain.Validado),
		"CV003": domain.NewCV("CV003", domain.Submetido),
	}}
}

func (r *InmemoryCVRepository) GetById(id string) (domain.CV, bool) {
	cv, ok := r.cvs[id]
	return cv, ok
}

func (r *InmemoryCVRepository) Save(cv domain.CV) error {
	r.cvs[cv.Id()] = cv
	return nil
}
