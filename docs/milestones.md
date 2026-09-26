Milestone 0: Histórico de puntuación consultable

Qué se entrega: un PMV que permite consultar el histórico de puntuación de cada corredor de la plantilla, reunido en un único sitio.
Objetivo: resolver que Alfonso ya no tenga que rebuscar entre hojas de cálculo de distintos años para valorar a un corredor.
Criterio de aceptación: dado un corredor conocido, se puede consultar su puntuación histórica por carrera y por año; para un corredor sin datos, se indica claramente que no hay histórico (no falla ni da un dato inventado).

Milestone 1: Alineación automática de convocatoria

Qué se entrega: un PMV que, dada una carrera, devuelve automáticamente los 7 corredores más adecuados para convocar, según su rendimiento histórico y descartando a los que están lesionados o de descanso.
Objetivo: resolver por completo que Alfonso tenga que cruzar a mano puntos históricos con la disponibilidad de cada corredor; la propuesta que reciba ya es directamente utilizable, sin revisión manual posterior.
Criterio de aceptación: dada una carrera conocida con corredores disponibles y con histórico suficiente, la alineación propuesta contiene exactamente 7 corredores, ninguno marcado como lesionado o en descanso, y son los de mejor rendimiento histórico entre los disponibles; si hay menos de 7 disponibles con datos, se indica explícitamente en vez de fallar o inventar corredores.
