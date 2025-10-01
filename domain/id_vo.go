package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type ID struct {
	value string
}

func NewID() (ID, error) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	newValue := fmt.Sprintf("Cand%03d", randomNum)

	return ID{value: newValue}, nil
}

func (i ID) Value() string {
	return i.value
}
