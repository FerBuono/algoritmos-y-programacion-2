package errors

import "fmt"

type ErrorReadingFile struct{}

func (e ErrorReadingFile) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParams struct{}

func (e ErrorParams) Error() string {
	return "ERROR: Faltan parámetros"
}

type DNIError struct{}

func (e DNIError) Error() string {
	return "ERROR: DNI incorrecto"
}

type DNIOutOfRegister struct{}

func (e DNIOutOfRegister) Error() string {
	return "ERROR: DNI fuera del padrón"
}

type EmptyQueue struct{}

func (e EmptyQueue) Error() string {
	return "ERROR: Fila vacía"
}

type ErrorVoterFraud struct {
	Dni int
}

func (e ErrorVoterFraud) Error() string {
	return fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", e.Dni)
}

type ErrorVoteType struct{}

func (e ErrorVoteType) Error() string {
	return "ERROR: Tipo de voto inválido"
}

type ErrorInvalidAlternative struct{}

func (e ErrorInvalidAlternative) Error() string {
	return "ERROR: Alternativa inválida"
}

type ErrorNoPreviousVotes struct{}

func (e ErrorNoPreviousVotes) Error() string {
	return "ERROR: Sin voto a deshacer"
}

type ErrorVotersWithoutVote struct{}

func (e ErrorVotersWithoutVote) Error() string {
	return "ERROR: Ciudadanos sin terminar de votar"
}
