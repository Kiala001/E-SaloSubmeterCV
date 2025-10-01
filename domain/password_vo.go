package domain

import "errors"

type Password struct {
	hash string
}

func NewPassword(hash string) (Password, error) {
	if hash == "" {
		return Password{}, errors.New("password hash cannot be empty")
	}

	if len(hash) < 6 {
		return Password{}, errors.New("password hash must be at least 6 characters long")
	}
	
	return Password{hash: hash}, nil
}

func (p Password) Hash() string {
	return p.hash
}