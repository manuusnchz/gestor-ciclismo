Milestone 0: Modelado del problema

Se entrega un PMV interno que consiste en usar Domain Driven Design para construir un lenguaje ubicuo, un vocabulario común que entiendan igual el cliente (Alfonso) y el desarrollador. En este milestone se trabaja únicamente con la primera historia de usuario, HU001, y el lenguaje se construye a partir de los conceptos que aparecen en ella. Todavía no hace nada, solo deja preparadas las estructuras de datos que se van a necesitar para programar la lógica del proyecto en el siguiente milestone. Se entrega como código dentro del repositorio, y con un documento en docs/ donde se explica cómo se ha analizado el problema.

Sabemos que es válido ya que en ese documento se ve que se ha aplicado la metodología sobre las historias de usuario, qué conceptos salen de ellas y qué se decide hacer con cada uno, y que todo lo que hay en el código viene de ese análisis y no de lo que se me haya ocurrido por mi cuenta.

Milestone 1: Lógica del problema

Se entrega un PMV interno con la primera parte de la lógica de HU001, programada sobre el modelo del milestone anterior: reconocer a un mismo corredor en las tablas de temporadas distintas, para que Alfonso pueda juntar sus puntos (#26). Solo se tienen en cuenta las temporadas que forman el histórico: la temporada en curso y las dos anteriores (#12). Junto con ella se entregan sus tests, las herramientas elegidas para escribirlos y ejecutarlos (#27) y la orden del gestor de tareas que los lanza. Se entrega como código y tests dentro del repositorio, todavía sin ninguna interfaz para el usuario.

Sabemos que es válido ya que, al lanzar los tests con el gestor de tareas, pasan todos estos:
- Un corredor que aparece en dos temporadas con el mismo texto se reconoce como el mismo.
- Un corredor al que le cambia la comunidad entre paréntesis de una temporada a otra se reconoce como el mismo.
- Un corredor que en una temporada tiene la comunidad entre paréntesis y en otra no se reconoce como el mismo.
- Dos corredores con distintos apellidos o nombre no se reconocen como el mismo.
- Un texto del ranking sin apellidos o sin nombre da error.
- Una edición de hace tres temporadas no aparece en el histórico.
- Una temporada en la que la carrera no se corrió aparece como sin edición.
- Un corredor que no está en la tabla de una edición aparece como que no participó.
- Un corredor que no participó se distingue de uno que participó y sacó 0 puntos.

Milestone 2: Histórico completo de HU001

Se entrega un PMV interno que completa la lógica de HU001 sobre lo entregado en el milestone anterior. Con el Milestone 1 ya se reconoce a un mismo corredor en las temporadas que forman el histórico que consulta Alfonso; en este se añade cómo se reconoce una misma carrera de un año a otro, para que pueda ver los puntos de un corredor en una carrera concreta a lo largo del histórico. También se añaden las comprobaciones automáticas del código y del fichero iv.yaml cada vez que se sube un cambio. Se entrega como código y tests dentro del repositorio, todavía sin ninguna interfaz para el usuario.

Sabemos que es válido ya que los tests de cada uno de sus issues pasan, también en los casos de error, y las comprobaciones automáticas se ejecutan al subir cada cambio sin dar fallos.
