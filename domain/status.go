package domain

type CVStatus string

const (
	Criado   CVStatus = "Criado"
	Validado CVStatus = "Validado"
	Submetido CVStatus = "Submetido"
)