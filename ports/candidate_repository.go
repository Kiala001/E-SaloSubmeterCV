package ports

import "esalo/domain"

type CandidateRepository interface {
	Length() int
	Save(Candidate domain.Candidate) error
	FindByEmail(email domain.Email) (domain.Candidate, bool)
}
