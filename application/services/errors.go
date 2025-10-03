package services

import "errors"

var (
	ErrCandidateAlreadyExists = errors.New("candidate with this email already exists")
	ErrCVNotFound		   = errors.New("cv not found")
)