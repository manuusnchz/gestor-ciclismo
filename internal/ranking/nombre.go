package ranking

type Nombre struct {
	nombre    string
	apellidos string
}

func NuevoNombre(nombre string, apellidos string) (Nombre, error) {
	if nombre == "" {
		return Nombre{}, ErrorNombreInvalido{Motivo: "falta el nombre"}
	}

	if apellidos == "" {
		return Nombre{}, ErrorNombreInvalido{Motivo: "faltan los apellidos"}
	}

	return Nombre{
		nombre:    nombre,
		apellidos: apellidos,
	}, nil
}
