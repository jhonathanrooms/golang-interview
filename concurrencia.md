# Concurrencia
Se refiere a la capacidad de un programa para gestionar múltiples tareas a la vez. No implica necesariamente que se estén ejecutando al mismo tiempo, sino que están progresando independientemente. En Go se implementa mediante las goroutines y channels, se pueden iniciar tareas concurrentes sin preocuparse directamente de la administración de threads.
## Goroutines 
Son el núcleo del modelo de concurrencia en Go, son funciones o métodos que se ejecutan de manera concurrente con otras goroutines. Se lanzan con la palabra clave "go" antes de la llamada a la función.
``` go
func myFunction() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go myFunction() // Lanza una nueva goroutine
    fmt.Println("Hello from main!")
}
```
### Características
* Ligeras y eficiente: Son mucho más ligeras que los hilos tradicionales, pues se administran en espacio de usuario en lugar de depender del kernel. Pueden lanzar miles de goroutines sin un gran costo en memoria y procesamiento.
* Administración de pila: Tienen una pila dinámica que puede crecer y reducirse según sea necesario, lo qUe permite manejar muchas más goroutines que hilos tradicionales.
* 
## Canales
Para que las goroutines se comuniquen entre sí y eviten condiciones de carrera, Go utirliza canales. Permiten el paso de mensajes entre goroutines de forma segura y se crean con make(chan type).
* Bloqueo: El envío y la recepción en un canal bloquean hasta que ambas operaciones estén listas. Esto permite la sincronización entre goroutines sin necesidad de bloqueos explícitos o variables de condición.
* Buffer: Es un área de almacenamiento temporal que se usa para guardar datos mientras se transfieren de un lugar a otro, se puede especificar una goroutine de la siguiente manera:
    * Sin buffer: Sin capacidad de almacenamiento, el receptor y el emisor deben sincronizarse.
    * Con buffer: Almacenan una cantidad limitada de valores sin bloquearse inmediatamente.
* Select: Permite esperar en múltiples canales, similar a una instrucción "switch" pero para operaciones de canal.
``` go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() { ch1 <- 42 }()
    go func() { ch2 <- 43 }()
    
    select {
    case val := <-ch1:
        fmt.Println("Received from ch1:", val)
    case val := <-ch2:
        fmt.Println("Received from ch2:", val)
    default:
        fmt.Println("No data received")
    }
}
```
## Context
Permite gestionar la duración de las goroutines, lo cual es esencial en aplicaciones complejas o con múltiples microservicios. Permiten cancelar o propagar cancelaciones en cascada para controlar el tiempo de vida de las goroutines.

## Paralelismo
Implica la ejecución de múltiples tareas al mismo tiempo, requiriendo un soporte a nivel de hardware. Go es concurrente por diseño, también puede ser paralelo si se ejecuta en un entorno con varios núcleos y se ajusta el número de hilos del sistema con runtime.GOMAXPROCS().
