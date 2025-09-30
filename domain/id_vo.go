package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type ID struct {
	value string
}

func (i ID) Value() string {
	return i.value
}

func (i ID) GenerateNew() (ID, error) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	newValue := fmt.Sprintf("Cand%03d", randomNum)

	if newValue == "" {
		return ID{}, errors.New("generated ID is empty")
	}

	return ID{value: newValue}, nil
}
