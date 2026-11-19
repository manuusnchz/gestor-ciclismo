# Herramientas de test

## Biblioteca de aserciones

### Criterios

1. **No escribir en el directorio del código**: al ejecutar los tests que la usan no puede crear ni modificar ficheros dentro del repositorio, o tiene que dejar configurar dónde lo hace, porque más adelante los tests se ejecutarán en un contenedor en el que no se puede escribir ahí. La que no lo cumpla queda fuera.
2. **Tiempo de ejecución**: cuánto tarda de media en ejecutarse el mismo test escrito con cada una, medido con hyperfine.

### Opciones

He buscado las opciones en la sección "Testing Frameworks" del repo awesome-go (https://github.com/avelino/awesome-go#testing-frameworks). Como ahí están mezclados los tres niveles con herramientas para otros tipos de test, me he quedado solo con las que sirven para comprobar el resultado de un test unitario. Además, considero no usar ninguna biblioteca, ya que el paquete `testing` permite comprobar los resultados sin ella. Las opciones son:

- Ninguna biblioteca: comprobar con un `if` y marcar el fallo con `t.Errorf` del paquete `testing`, que es lo que recomienda la wiki de Go (https://go.dev/wiki/TestComments).
- go-cmp (https://github.com/google/go-cmp): lo mantiene el equipo de Go. Solo compara dos valores y devuelve la diferencia; el fallo se marca igualmente con `t.Errorf`.
- Bibliotecas con funciones de aserción: assert (https://github.com/go-playground/assert), be (https://github.com/earthboundkid/be), go-testdeep (https://github.com/maxatome/go-testdeep), is (https://github.com/matryer/is), should (https://github.com/Kairum-Labs/should), trial (https://github.com/jgroeneveld/trial), Tt (https://github.com/vcaesar/tt), testza (https://github.com/MarvinJWendt/testza), got (https://github.com/ysmood/got) y los paquetes `assert` de Testify (https://github.com/stretchr/testify) y de gotest.tools (https://github.com/gotestyourself/gotest.tools). testza y got se presentan como frameworks, pero lo que añaden a `testing` son aserciones.
- Bibliotecas de matchers o de estilo BDD, que se leen más como una frase: gomega (https://github.com/onsi/gomega), gocrest (https://github.com/corbym/gocrest), Hamcrest (https://github.com/rdrdr/hamcrest) y go-testpredicate (https://github.com/maargenton/go-testpredicate).

## Test runner o framework

### Criterios

1. **No escribir en el directorio del código**: igual que en las aserciones, al ejecutar los tests no puede crear ni modificar ficheros dentro del repositorio, o tiene que dejar configurar dónde lo hace. El que no lo cumpla queda fuera.
2. **Tiempo de ejecución**: cuánto tarda de media en ejecutarse el mismo test escrito con cada uno, medido con hyperfine.

## Herramienta CLI para ejecutar los tests

### Criterios

1. **No escribir en el directorio del código**: al lanzar los tests no puede crear ni modificar ficheros dentro del repositorio, o tiene que dejar configurar dónde lo hace. La que no lo cumpla queda fuera.
2. **Tiempo de ejecución**: cuánto tarda de media en lanzar los tests, medido con hyperfine.
