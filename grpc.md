# gRPC
Es un marco de comunicación remota que permite a los servicios interactuar de manera eficiente. Fue desarrollado por Google y utiliza HTTP/2 para transmitir datos en formato binario (Protocol Buffers) en lugar de texto, lo que mejora la eficiencia en comparación con otras alternativas como REST. Es especialmente útil en sistemas distribuidos y microservicios porque permiten comunicación rápida y de bajo consumo de recursos, además de proporcionar soporte nativo para varios lenguajes, incluida la generación automática de código de cliente y servidor.

## Conceptos Básicos
1. Protocol Buffers: Es un lenguaje de serialización binario que gRPC usa para definir los mensajes y servicios. Es más eficiente que JSON o XML porque reduce el tamaño de los datos y mejora la velocidad de transmisión. Los mensajes y servicios se definen en un archivo **.proto**.
2. Servicios y métodos: Un servicio es una colección de métodos remotos que un cliente puede llamar. Cada método define una entrada y una salida, que son mensajes definidos en el archivo **.proto**.
3. Tipos de comunicación gRPC: Soporta cuatro tipos de métodos:
    * Unary: Cliente envía una solicutud y recibe una respuesta (similar a REST).
    * Server streaming: Cliente envía una solicitud y recibe una secuencia de respuestas.
    * Client streaming: Cliente envía una secuencia de solicitudes y recibe una sola respuesta.
    * Bidirectional streaming: Cliente y servidor envían secuencias de mensajes.