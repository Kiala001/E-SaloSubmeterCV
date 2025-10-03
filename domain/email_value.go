package domain

import (
	"regexp"
)

type Email struct {
	address string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(address string) (Email, error) {
	if address == "" {
		return Email{}, ErrEmailCannotBeEmpty
	}
	if !emailRegex.MatchString(address) {
		return Email{}, ErrInvalidEmailFormat
	}

	return Email{address: address}, nil
}

func (e Email) String() string {
	return e.address
}
