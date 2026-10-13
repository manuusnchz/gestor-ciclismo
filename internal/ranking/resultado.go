package ranking

type Resultado struct {
	participo bool
	puntos    int
}

func ResultadoNoParticipo() Resultado {
	return Resultado{}
}

func ResultadoParticipo(puntos int) (Resultado, error) {
	if puntos < 0 {
		return Resultado{}, ErrorPuntosNegativos{Valor: puntos}
	}

	return Resultado{
		participo: true,
		puntos:    puntos,
	}, nil
}
