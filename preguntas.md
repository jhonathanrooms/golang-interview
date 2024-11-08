# Respuestas a Preguntas Técnicas
## ¿Cómo funciona el garbage collector en Go y cómo impacta el rendimiento de una aplicación?
El garbage collector en Go es concurrente y detiene el mundo ("stop-the-world") de manera mínima. Recolecta memoria no utilizada para liberar recursos, pero puede causar latencia en aplicaciones de alta concurrencia.
## Explica una situación donde usarías un canal sin buffer en lugar de uno con buffer.
Usaría un canal sin buffer para sincronizar el inicio de una goroutine y asegurar que una función no se ejecute hasta que la otra termine, creando una comunicación de "mano a mano".
## ¿Cómo manejarías un deadlock en un sistema concurrente en Go?
Evitaría deadlocks diseñando cuidadosamente el orden de bloqueos y usando canales para coordinar goroutines. También podría usar select con default para evitar bloqueos.
## ¿Qué patrones o principios usas para estructurar proyectos en Go? ¿Cómo organizarías un servicio en una arquitectura de microservicios?
Siguiendo Clean Architecture, organizaría los paquetes en domain (entidades), usecase (lógica de negocio), repository (persistencia), y api (controladores y servicios).
## ¿Cómo evitarías las condiciones de carrera (race conditions) en una aplicación con concurrencia?
Utilizaría sync.Mutex para bloquear acceso a secciones críticas o canales para sincronización segura entre goroutines.
## ¿Qué ventajas y desventajas tiene Go comparado con otros lenguajes para aplicaciones de alto rendimiento?
Ventajas: ejecución rápida, concurrencia nativa con goroutines. Desventajas: garbage collector puede afectar latencia en algunos casos, y falta de soporte para programación genérica (aunque ya está mejorando en versiones recientes).
## ¿Cómo implementarías una solución resiliente que maneje fallos de red o de servicios externos en Go?
Usaría context.Context para manejar timeouts y cancelaciones, y patrones de reintento con backoff exponencial.
## ¿Cuándo es más útil una WaitGroup versus un canal para sincronizar goroutines?
WaitGroup es útil para esperar a que un conjunto de goroutines terminen, mientras que los canales son mejores para la comunicación y sincronización de datos entre goroutines.
## ¿Cómo gestionarías la inyección de dependencias en Go sin frameworks específicos?
Usaría interfaces y asignación explícita de dependencias en el constructor, permitiendo inyección de mocks durante pruebas.
## ¿Cómo asegurarías que tu código es eficiente en cuanto a memoria y procesamiento?
Usaría herramientas de profiling (pprof), optimizaría estructuras de datos, evitaría copias innecesarias y mediría los resultados para asegurar eficiencia.