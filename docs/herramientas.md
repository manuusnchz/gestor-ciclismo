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
