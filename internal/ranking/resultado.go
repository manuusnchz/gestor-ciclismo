package ranking

type Resultado struct {
	participo bool
	puntos    Puntos
}

func ResultadoNoParticipo() Resultado {
	return Resultado{}
}

func ResultadoParticipo(puntos Puntos) Resultado {
	return Resultado{
		participo: true,
		puntos:    puntos,
	}
}
