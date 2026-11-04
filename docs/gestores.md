# Herramientas de automatización

## Gestor de dependencias

Según la web oficial de Go, en las notas de Go 1.16 (https://go.dev/doc/go1.16#modules), desde Go 1.16 `go` trabaja con los módulos de Go por defecto. Además, en la referencia de los módulos (https://go.dev/ref/mod) se ve que el fichero `go.mod` guarda la ruta del módulo, la versión de Go y las dependencias, por lo que la versión de Go queda en un único sitio. Por otro lado, no hay que instalar nada aparte, ya que viene dentro de `go`. Esto es simplemente una facilidad y no es lo que me hace decidirme, pero lo compruebo en la prueba.

Otras opciones usadas anteriormente son dep (https://github.com/golang/dep), que está archivado desde 2020; glide (https://github.com/Masterminds/glide), que en su README pone textualmente "The Go community is now using Go Modules to handle dependencies. Please consider using that instead of Glide. Glide is now mostly unmaintained."; govendor (https://github.com/kardianos/govendor), que está archivado y aparte en su README pone "Use Go modules"; y godep (https://github.com/tools/godep), que también está archivado. También está el modo GOPATH, que según la documentación de `go` (https://pkg.go.dev/cmd/go) sigue existiendo como modo heredado (con `GO111MODULE=off`), pero no gestiona versiones y no es el modo por defecto. Por eso no hago la prueba con ninguna de ellas: están archivadas, indican directamente que se usen los módulos de Go o, en el caso del modo GOPATH, no gestionan versiones.

### Prueba

Para comprobarlo he usado un contenedor Docker con la imagen oficial de Go, con la misma versión que pide `go.mod` (`golang:1.27.3`). Así parto de una máquina limpia que solo tiene Go. El repositorio se monta en solo lectura y cada ejecución se hace en un contenedor nuevo (`--rm`), para que una no influya en las mediciones de los otros.

- `time go mod download`, tres veces, cada una en un contenedor nuevo: descarga las dependencias que declara `go.mod` y mide cuánto tarda. Las tres veces responde "no module dependencies to download" y tarda 0,008 s, 0,007 s y 0,007 s, ya que en mi proyecto por ahora no hay dependencias que descargar

![Tiempos de go mod download](../imgs/prueba1.png)

- `go version`: muestra que Go 1.27.3 ya viene en la imagen y que no he tenido que instalar nada más para gestionar las dependencias.
- `go list -m all`: lista todos los módulos que usa el proyecto. Solo aparece el propio, por que no hay ninguna dependencia externa.
- `go mod verify`: comprueba que los módulos descargados no se han modificado. Responde "all modules verified".

![go version, go list -m all y go mod verify](../imgs/prueba2.png)

Ahora mismo instalar las dependencias no cuesta prácticamente nada (unos 0,007 s), porque el proyecto solo usa la biblioteca estándar y no tiene ninguna dependencia externa. Tampoco ha hecho falta instalar nada aparte de Go: en el contenedor solo estaba la imagen oficial de Go y todo ha funcionado. Por eso la tarea `install` (`go mod download`) hoy no hace nada; empezará a descargar algo cuando se añada alguna dependencia a `go.mod`.
