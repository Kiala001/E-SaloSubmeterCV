package domain

import (
	"regexp"
	"strings"
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return Name{}, ErrNameCannotBeEmpty
	}

	if matched, _ := regexp.MatchString(`[0-9]`, value); matched {
		return Name{}, ErrNameCannotContainNumbers
	}

	parts := strings.Fields(value)
	if len(parts) < 2 {
		return Name{}, ErrNameShouldContainAtLeastFirstAndLast
	}

	return Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}
