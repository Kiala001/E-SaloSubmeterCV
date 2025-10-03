package domain

type CVStatus string

const (
	CRIADO   CVStatus = "Criado"
	VALIDADO CVStatus = "Validado"
	SUBMETIDO CVStatus = "Submetido"
)