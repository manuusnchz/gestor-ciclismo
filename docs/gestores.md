# Herramientas de automatización

## Gestor de dependencias

### Criterios

Para elegir sigo los criterios generales señalados en el guión, en el que se dice que "Si hay un estándar, primero hay que conocerlo; a continuación simplemente hay que seguirlo y adaptar el resto de los criterios al mismo". Los dos primeros son obligatorios, y entre las herramientas que los cumplan la elección la hacen los dos últimos.

1. **Estándar**: que sea el gestor de dependencias que define el propio Go en su documentación oficial (https://go.dev/ref/mod).
2. **Funciona con la versión del proyecto**: que se pueda usar con Go 1.27.3.
3. **Coste de instalación**: qué hay que instalar aparte de Go y cuánto tarda, medido en un contenedor limpio con la imagen oficial.
4. **Coste de ejecución**: cuánto tarda la orden que instala las dependencias del proyecto, medido en ese mismo contenedor.

### Herramientas consideradas

- **Módulos de Go** (https://go.dev/ref/mod):
  - Cumple los dos criterios obligatorios ya que es estandar y desde la version 1.16 de Go, este trabaja con los módulos por defecto, y por supuesto cumple el segundo criterio de funcionar con Go 1.27.3. Los siguientes dos criterios se medirán más adelante.

Al buscar gestores de dependencias para Go también aparecen estas herramientas, que se usaban antes de los módulos y siguen saliendo en tutoriales y proyectos antiguos. Ninguna cumple el primer criterio.

- [Modo GOPATH](https://pkg.go.dev/cmd/go): no es un gestor aparte, sino un modo de `go`. La documentación de `go` lo trata como un modo heredado, y desde Go 1.16 el modo por defecto son los módulos.
- [dep](https://github.com/golang/dep): no es el estándar. Está archivado y su README dice "As of 2020, Dep is deprecated and archived in favor of Go modules".
- [glide](https://github.com/Masterminds/glide): no es el estándar. Su README dice "The Go community is now using Go Modules to handle dependencies. Please consider using that instead of Glide."
- [govendor](https://github.com/kardianos/govendor): no es el estándar. Está archivado y su README dice "Use Go modules".
- [godep](https://github.com/tools/godep): no es el estándar. Está archivado y su README dice "Please use dep or another tool instead", y dep, a su vez, remite a los módulos de Go.

Por tanto, la única herramienta que cumple los dos criterios obligatorios y que pruebo es la de los módulos de Go.

### Prueba

Para medir el coste de instalación y el de ejecución he usado un contenedor Docker con la imagen oficial de Go, con la misma versión que fija `go.mod` (`golang:1.27.3`). Así parto de una máquina limpia que solo tiene Go.

- `go version`: muestra que Go 1.27.3 ya viene en la imagen, así que los módulos funcionan con la versión del proyecto y no hay que instalar nada aparte de Go. El coste de instalación es cero.
- `go list -m all`: lista todos los módulos que usa el proyecto. Solo aparece el propio, porque no hay ninguna dependencia externa.
- `go mod verify`: comprueba que los módulos descargados no se han modificado. Responde "all modules verified".

![go version, go list -m all y go mod verify](../imgs/prueba2.png)

- `time go mod download`, tres veces, cada una en un contenedor nuevo, para medir el coste de ejecución: es la orden que instala las dependencias que declara `go.mod`. Las tres veces responde "no module dependencies to download" y tarda 0,008 s, 0,007 s y 0,007 s.

![Tiempos de go mod download](../imgs/prueba1.png)

Ahora mismo instalar las dependencias no cuesta prácticamente nada porque el proyecto solo usa la biblioteca estándar. Este tiempo cambiará cuando se añada alguna dependencia a `go.mod`.

### Decisión final

Los módulos de Go son la única herramienta que cumple los dos criterios obligatorios: son el estándar de Go y funcionan con la versión del proyecto. Además, no hay que instalar nada aparte de Go y ahora mismo instalar las dependencias tarda unos 0,007 s. Por eso elijo los módulos de Go como gestor de dependencias.
