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

- `time go mod download`, tres veces, cada una en un contenedor nuevo: es la orden que instala las dependencias que declara `go.mod`. Las tres veces responde "no module dependencies to download" y tarda 0,008 s, 0,007 s y 0,007 s, porque no tiene nada que descargar.

![Tiempos de go mod download](../imgs/prueba1.png)

Por tanto, esta medida no sirve todavía para saber cuánto cuesta instalar dependencias con los módulos de Go: el proyecto solo usa la biblioteca estándar y la orden no tiene nada que descargar.

### Decisión final

Los módulos de Go son la única herramienta que cumple los dos criterios obligatorios: son el estándar de Go y funcionan con la versión del proyecto. Además, no hay que instalar nada aparte de Go y, por ahora, el proyecto no tiene dependencias que descargar. Por eso elijo los módulos de Go como gestor de dependencias.

## Gestor de tareas

### Criterios

1. **Tiempo de instalación**: qué hay que instalar aparte de Go y cuánto tarda.
2. **Tiempo de ejecución**: cuánto tarda en lanzar la tarea `check`.

### Opciones

He buscado las opciones en "Build Automation" dentro del repo awesome-go (https://github.com/avelino/awesome-go#build-automation). Como en esa sección no todo lo que hay son gestores de tareas me he quedado solo con las que sirven para escribir tareas en un fichero y lanzarlas. A esas les he sumado just (https://github.com/casey/just), que no aparece en la lista porque no está escrita en Go, pero sirve para cualquier lenguaje, y make.

De esa sección quedan fuera air, anko, gaper y realize, que sirven para recargar la aplicación mientras se desarrolla, y gob, cuyo repositorio ya no existe. mmake no lo pruebo porque, según su README, funciona por encima de make, así que no es una alternativa a make.

Pruebo make, Task, just, Mage, 1build, taskctl, xc, goyek, rex y gilbert.

### Pruebas

#### make

make no hay que instalarlo, ya que lo tengo instalado (GNU Make 4.4.1). `make check` ha tardado 0,036 s, 0,014 s y 0,015 s.

![make](../imgs/make.png)

#### Task

Task (https://github.com/go-task/task) hay que instalarlo. Con el script de su documentación oficial (https://taskfile.dev/installation/), la instalación ha tardado 5,603 s y ha instalado la versión 3.54.0. `task check` ha tardado 0,059 s, 0,067 s y 0,062 s.

![Task](../imgs/task.png)

#### just

just (https://github.com/casey/just) hay que instalarlo. Con el script del apartado "Pre-built Binaries" de su README, la instalación ha tardado 6,317 s y ha instalado la versión 1.59.0. `just check` ha tardado 0,013 s, 0,015 s y 0,017 s.

![just](../imgs/just.png)

#### Mage

Mage (https://github.com/magefile/mage) hay que instalarlo. Las tareas se escriben como funciones de Go en un fichero `magefile.go`. Con `go install`, la instalación ha tardado 3,137 s y ha instalado la versión 1.17.2. `mage check` ha tardado 0,240 s, 0,120 s y 0,118 s; la primera vez tarda más porque compila antes el `magefile.go`.

![Mage](../imgs/mage.png)

#### 1build

1build (https://github.com/gopinath-langote/1build) hay que instalarlo. Descargando el programa ya compilado de su última versión publicada, como indica su README, la instalación ha tardado 1,514 s y ha instalado la versión 1.10.4. `1build check` ha tardado 0,014 s, 0,016 s y 0,016 s.

![1build](../imgs/1build.png)

#### taskctl

taskctl (https://github.com/taskctl/taskctl) hay que instalarlo. El método para Linux de su README da un error 404, porque el fichero que indica no existe en su última versión, así que lo he instalado con el paquete `.deb`, que es la otra forma que da el README. La instalación ha tardado 1,041 s y ha instalado la versión 1.4.2. `taskctl run check` ha tardado 0,017 s, 0,015 s y 0,024 s.

![taskctl](../imgs/taskctl.png)

#### xc

xc (https://github.com/joerdav/xc) hay que instalarlo. Las tareas se escriben en un fichero Markdown, bajo un título `## Tasks`. Descargando el programa ya compilado, como indica su guía oficial (https://xcfile.dev/getting-started/#installation), la instalación ha tardado 0,916 s y ha instalado la versión 0.9.0. `xc -f tareas.md check` ha tardado 0,018 s, 0,017 s y 0,017 s.

![xc](../imgs/xc.png)

#### goyek

goyek (https://github.com/goyek/goyek) no es un programa, sino una biblioteca de Go: las tareas se escriben en Go en una carpeta `build/` y se lanzan con `go run`. Siguiendo su README, la instalación es `go mod tidy`, que ha tardado 3,847 s y ha descargado la versión 3.0.1. `go run . check` ha tardado 0,727 s, 0,087 s y 0,102 s; la primera vez tarda más porque compila antes la carpeta `build/`.

![goyek](../imgs/goyek.png)

#### rex

rex (https://github.com/rexrun-dev/rex) no viene en la imagen. Lo he instalado con `go install` , igual que con Mage, que ha tardado 8,375 s en descargarlo y compilarlo, y ha instalado la versión 0.5.0. Las tareas se escriben en un fichero `rex.toml`, en la sección `[commands]`. Según su README, si el proyecto ya tiene un Makefile, rex lo usa antes que su propio fichero, así que en la copia del contenedor he borrado el Makefile para probar de verdad rex y no make. Las dos tareas se han ejecutado correctamente desde `rex.toml` y `go.mod` no ha cambiado, por lo que rex cumple los dos primeros criterios, pero es la herramienta que más ha tardado en instalarse de todas las que he probado.

![Prueba de rex](../imgs/prueba-rex.png)

Para medir la ejecución he lanzado `rex check` tres veces en otro contenedor limpio, también sin el Makefile, y ha tardado 0,011 s, 0,009 s y 0,007 s, prácticamente lo mismo que lanzar `gofmt` directamente.

![Ejecución de check con rex](../imgs/ejecucion-rex.png)

### Decisión final

Después de las pruebas , descarto `go` porque no cumple el criterio 1, taskctl porque no se puede instalar con el método de su documentación, y goyek porque añade dependencias a `go.mod` (criterio 2). make, Task, just, Mage, 1build, xc y rex cumplen los dos criterios obligatorios y todas tienen commits en el último año, así que tampoco las separa la deuda técnica (criterio 4). Lo que decide es el criterio 5: make ya viene en la imagen oficial de Go, así que no hay que instalar nada aparte, mientras que las demás tardan en instalarse entre los 0,584 s de xc y los 8,375 s de rex. Al lanzar `check`, make tarda lo mismo que `gofmt` directamente, mientras que Task tarda unos 80 ms más cada vez y Mage unos 370 ms más la primera vez, porque tiene que compilar el `magefile.go`. Por eso elijo make, con un Makefile en la raíz del repositorio con las tareas `install` y `check`.
