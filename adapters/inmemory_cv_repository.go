package adapters

import "esalo/domain"

type InmemoryCVRepository struct {
	CVs map[string]domain.CV
}

func NewInmemoryCVRepository() *InmemoryCVRepository {
	return &InmemoryCVRepository{CVs: map[string]domain.CV{
		"CV001": {Id: "CV001", Status: ""},
		"CV002": {Id: "CV002", Status: "Validado"},
		"CV003": {Id: "CV003", Status: "Submetido"},
	}}
}

func (r *InmemoryCVRepository) GetById(id string) (domain.CV, bool) {
	cv, ok := r.CVs[id]
	return cv, ok
}

func (r *InmemoryCVRepository) Save(cv domain.CV) error {
	r.CVs[cv.Id] = cv
	return nil
}
