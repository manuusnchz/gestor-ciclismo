# Herramientas del proyecto

## Gestor de dependencias

El repositorio no tenía nada que dijera qué hace falta para compilar el
código: ni el módulo, ni las dependencias, ni la versión de Go (#23).

Para elegir me he fijado en que sea lo estándar del lenguaje, que siga
mantenido y que no haya que instalar nada aparte de Go, ya que ahora mismo
el proyecto solo usa la biblioteca estándar.

En Go las opciones que hay son los módulos de Go, dep, glide, govendor,
godep y el modo GOPATH de antes. Haciendo una búsqueda para ver cual era el 
más adecuado para usar, he descubierto que menos los módulos de Go, los demás
están completamente obsoletos. Debido a esto usaré los módulos de Go , que vienen
integrados en el propio Go. 

El fichero [`go.mod`](../go.mod) declara el módulo y la versión de Go con la
que se compila, así que la versión está en un único sitio, como se habló
en #6.

## Gestor de tareas

Para instalar lo que necesita el proyecto o comprobar que compila hay que
saber qué órdenes de Go escribir, y no hay un sitio donde estén todas
(#24). Además `go` no tiene ningún fichero donde poner las tareas ni una
orden `check`, por lo que cada vez hay que acordarse de la orden y
escribirla a mano.

Antes de mirar herramientas he pensado qué necesito: las tareas son pocas
(instalar, comprobar que compila y más adelante los tests) y todas son
lanzar una orden de `go`, así que no necesito escribirlas en un lenguaje de
programación. Tampoco quiero añadir dependencias al proyecto, que ahora
solo usa la biblioteca estándar, y tiene que seguir mantenido.

Las opciones que he visto son make, Task, Mage, goyek, just, Earthly y
sistemas más grandes como Bazel o Please. Mage y goyek escriben las tareas
en Go y entran como dependencia en `go.mod`. Task y just han sacado
versión este mismo mes (v3.54.0 y 1.59.0) así que están mantenidos, pero
hay que instalarlos a parte y cada uno usa su propio fichero. Earthly ya
no se mantiene y Bazel o Please son demasiado para un proyecto con un solo
paquete.

Debido a esto uso make, que lanza las órdenes tal cual y no añade nada al
proyecto. Su última versión es la 4.4.1, de 2023, y además es la única de
todas estas que está dentro del estándar POSIX
([IEEE Std 1003.1-2024](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)),
por lo que el `Makefile` no depende de un programa en concreto. Lo malo es
que no siempre viene instalado, pero si no está se instala con el gestor
de paquetes del sistema.
