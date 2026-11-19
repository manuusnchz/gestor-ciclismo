# Herramientas de test

Los tests del Milestone 1 comprueban la consulta del histórico de un corredor en una carrera. La consulta recibe el nombre del corredor y la temporada en curso, y devuelve lo que hizo en esa temporada y en las dos anteriores: si la carrera no se corrió ese año, si no participó o cuántos puntos sacó, aunque sean 0. Por eso no me basta con comprobar un sí o un no, tengo que comparar lo que devuelve para las tres temporadas con lo esperado. Varios casos son el mismo test con distintos datos, y hay uno de error: al cargar una edición con un texto del ranking sin apellidos o sin nombre, hay que comprobar que el error es `ErrorNombreInvalido` y no uno cualquiera.

## Biblioteca de aserciones

### Criterios

1. **Comparar lo que devuelve la consulta**: tiene que poder comparar lo que devuelve la consulta para las tres temporadas con lo esperado, y comprobar que un error es de un tipo concreto, como `ErrorNombreInvalido`. La que no lo permita queda fuera.
2. **Tiempo de ejecución**: cuánto tarda de media en ejecutarse el mismo test escrito con cada una, medido con hyperfine.

## Test runner o framework

### Criterios

1. **Repetir un test con distintos datos**: tiene que permitir escribir el test una vez y lanzarlo con cada caso. El que no lo permita queda fuera.
2. **Tiempo de ejecución**: cuánto tarda de media en ejecutarse el mismo test escrito con cada uno, medido con hyperfine.

## Herramienta CLI para ejecutar los tests

### Criterios

1. **No escribir en el directorio del código**: al lanzar los tests no puede crear ni modificar ficheros dentro del repositorio, o tiene que dejar configurar dónde lo hace, porque más adelante los tests se ejecutarán en un contenedor en el que no se puede escribir ahí. El que no lo cumpla queda fuera.
2. **Tiempo de ejecución**: cuánto tarda de media en lanzar los tests, medido con hyperfine.
