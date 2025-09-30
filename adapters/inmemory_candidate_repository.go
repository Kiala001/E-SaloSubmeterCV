package adapters

import "esalo/domain"

type InmemorycandidateRepository struct {
	Candidates map[string]domain.Candidate
}

func NewInmemorycandidateRepository() *InmemorycandidateRepository {
	return &InmemorycandidateRepository{
		Candidates: make(map[string]domain.Candidate),
	}
}

func (r *InmemorycandidateRepository) Length() int {
	return len(r.Candidates)
}

func (r *InmemorycandidateRepository) Save(Candidate domain.Candidate) error {
	r.Candidates[Candidate.Id] = Candidate
	return nil
}
