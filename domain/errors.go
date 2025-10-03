package domain

import "errors"

var (
	ErrCVIDCannotBeEmpty    = errors.New("cv id cannot be empty")
	
	ErrCVAlreadyValidated = errors.New("cv already validated")
	ErrCannotValidateCV    = errors.New("cannot validate a cv that is already submitted")
	
	ErrCVAlreadySubmitted = errors.New("cv already submitted")
	ErrCannotSubmitCV      = errors.New("cannot submit a cv that is not validated")
	
	ErrNameCannotBeEmpty    = errors.New("name cannot be empty")
	ErrNameCannotContainNumbers = errors.New("name cannot contain numbers")
	ErrNameShouldContainAtLeastFirstAndLast = errors.New("name should contain at least first name and last name")

	ErrEmailCannotBeEmpty   = errors.New("email cannot be empty")
	ErrInvalidEmailFormat   = errors.New("invalid email format")

	ErrPasswordCannotBeEmpty = errors.New("password cannot be empty")
	ErrPasswordTooShort      = errors.New("password must be at least 6 characters long")

)