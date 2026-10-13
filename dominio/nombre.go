package dominio

import (
	"regexp"
	"strings"
)

var parentesisRegex = regexp.MustCompile(`\([^)]*\)`)
var espaciosRegex = regexp.MustCompile(`\s+`)

type Nombre struct {
	valor string
}

func NuevoNombre(raw string) (Nombre, error) {
	limpio := parentesisRegex.ReplaceAllString(raw, " ")
	limpio = quitarAcentos(limpio)
	limpio = espaciosRegex.ReplaceAllString(limpio, " ")
	limpio = strings.TrimSpace(limpio)

	if limpio == "" {
		return Nombre{}, ErrorNombreInvalido{
			Motivo: "el nombre no puede estar vacío",
			Texto:  raw,
		}
	}

	if strings.Contains(limpio, ",") {
		partes := strings.Split(limpio, ",")
		if len(partes) != 2 {
			return Nombre{}, ErrorNombreInvalido{
				Motivo: "formato con coma inválido",
				Texto:  raw,
			}
		}

		apellidos := strings.TrimSpace(espaciosRegex.ReplaceAllString(partes[0], " "))
		nombre := strings.TrimSpace(espaciosRegex.ReplaceAllString(partes[1], " "))
		if apellidos == "" || nombre == "" {
			return Nombre{}, ErrorNombreInvalido{
				Motivo: "deben existir apellidos y nombre a ambos lados de la coma",
				Texto:  raw,
			}
		}

		limpio = apellidos + ", " + nombre
	} else {
		palabras := strings.Fields(limpio)
		if len(palabras) < 2 {
			return Nombre{}, ErrorNombreInvalido{
				Motivo: "debe incluir al menos nombre y apellidos",
				Texto:  raw,
			}
		}

		limpio = strings.Join(palabras, " ")
	}

	return Nombre{
		valor: strings.ToUpper(limpio),
	}, nil
}

func quitarAcentos(s string) string {
	var resultado strings.Builder
	resultado.Grow(len(s))

	for _, caracter := range s {
		switch caracter {
		case 'Á', 'À', 'Â', 'Ä':
			resultado.WriteRune('A')
		case 'á', 'à', 'â', 'ä':
			resultado.WriteRune('a')
		case 'É', 'È', 'Ê', 'Ë':
			resultado.WriteRune('E')
		case 'é', 'è', 'ê', 'ë':
			resultado.WriteRune('e')
		case 'Í', 'Ì', 'Î', 'Ï':
			resultado.WriteRune('I')
		case 'í', 'ì', 'î', 'ï':
			resultado.WriteRune('i')
		case 'Ó', 'Ò', 'Ô', 'Ö':
			resultado.WriteRune('O')
		case 'ó', 'ò', 'ô', 'ö':
			resultado.WriteRune('o')
		case 'Ú', 'Ù', 'Û', 'Ü':
			resultado.WriteRune('U')
		case 'ú', 'ù', 'û', 'ü':
			resultado.WriteRune('u')
		default:
			resultado.WriteRune(caracter)
		}
	}

	return resultado.String()
}
