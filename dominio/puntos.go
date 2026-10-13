package dominio

type Puntos struct {
	valor int
}

func NuevosPuntos(valor int) (Puntos, error) {
	if valor < 0 {
		return Puntos{}, ErrorPuntosNegativos{
			Valor: valor,
		}
	}

	return Puntos{
		valor: valor,
	}, nil
}
