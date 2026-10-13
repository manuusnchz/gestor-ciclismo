package dominio

import "time"

type Temporada struct {
	anio int
}

func NuevaTemporada(anio int) (Temporada, error) {
	if anio <= 0 {
		return Temporada{}, ErrorTemporadaInvalida{
			Anio: anio,
		}
	}

	anioActual := time.Now().Year()
	if anio > anioActual {
		return Temporada{}, ErrorTemporadaNoComenzada{
			Anio:       anio,
			AnioActual: anioActual,
		}
	}

	return Temporada{
		anio: anio,
	}, nil
}
