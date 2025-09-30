package domain

import (
	"errors"
	"regexp"
	"strings"
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return Name{}, errors.New("name cannot be empty or only spaces")
	}

	if matched, _ := regexp.MatchString(`[0-9]`, value); matched {
		return Name{}, errors.New("name cannot contain numbers")
	}

	parts := strings.Fields(value)
	if len(parts) < 2 {
		return Name{}, errors.New("name must contain at least first name and last name")
	}

	return Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}