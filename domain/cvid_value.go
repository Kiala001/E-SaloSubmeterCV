package domain

type CVId struct {
    value string
}

func NewCVId(value string) (CVId, error) {
    if value == "" {
        return CVId{}, ErrCVIDCannotBeEmpty
    }
    return CVId{value: value}, nil
}

func (id CVId) Value() string {
    return id.value
}
