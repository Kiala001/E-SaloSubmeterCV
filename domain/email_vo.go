package domain

import (
	"errors"
	"regexp"
)

type Email struct {
	address string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(address string) (Email, error) {
	if address == "" {
		return Email{}, errors.New("email address cannot be empty")
	}
	if !emailRegex.MatchString(address) {
		return Email{}, errors.New("invalid email address")
	}

	return Email{address: address}, nil
}

func (e Email) String() string {
	return e.address
}
