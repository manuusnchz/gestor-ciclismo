package ranking

import "fmt"

type ErrorPuntosNegativos struct {
	Valor int
}

func (e ErrorPuntosNegativos) Error() string {
	return fmt.Sprintf("los puntos no pueden ser negativos: %d", e.Valor)
}

type ErrorNombreInvalido struct {
	Motivo string
}

func (e ErrorNombreInvalido) Error() string {
	return fmt.Sprintf("nombre inválido: %s", e.Motivo)
}
