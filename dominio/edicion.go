package dominio

type Edicion struct {
	NombrePublicado string
	Resultados      map[Nombre]Resultado
}
