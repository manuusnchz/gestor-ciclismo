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

type ErrorTemporadaInvalida struct {
	Anio int
}

func (e ErrorTemporadaInvalida) Error() string {
	return fmt.Sprintf("temporada inválida: %d", e.Anio)
}

type ErrorTemporadaNoComenzada struct {
	Anio       int
	AnioActual int
}

func (e ErrorTemporadaNoComenzada) Error() string {
	return fmt.Sprintf("la temporada %d aún no ha comenzado; el año actual es %d", e.Anio, e.AnioActual)
}
