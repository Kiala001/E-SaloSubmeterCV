package domain

type Password struct {
	hash string
}

func NewPassword(hash string) (Password, error) {
	if hash == "" {
		return Password{}, ErrPasswordCannotBeEmpty
	}

	if len(hash) < 6 {
		return Password{}, ErrPasswordTooShort
	}

	return Password{hash: hash}, nil
}

func (p Password) Hash() string {
	return p.hash
}
