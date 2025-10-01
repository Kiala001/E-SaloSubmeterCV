package adapters

import "esalo/domain"

type InmemoryCandidateRepository struct {
	candidates map[string]domain.Candidate
}

func NewInmemoryCandidateRepository() *InmemoryCandidateRepository {
	return &InmemoryCandidateRepository{
		candidates: make(map[string]domain.Candidate),
	}
}

func (r *InmemoryCandidateRepository) Length() int {
	return len(r.candidates)
}

func (r *InmemoryCandidateRepository) Save(Candidate domain.Candidate) error {
	r.candidates[Candidate.ID().Value()] = Candidate
	return nil
}

func (r *InmemoryCandidateRepository) FindByEmail(email domain.Email) (domain.Candidate, bool) {
	for _, candidate := range r.candidates {
		if candidate.Email() == email {
			return candidate, true
		}
	}
	return domain.Candidate{}, false
}