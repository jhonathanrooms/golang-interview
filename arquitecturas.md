# Arquitecturas de Software
## Monolítica
Toda la lógica de la aplicación se encuentra dentro de un solo paquete o aplicación desplegada, este enfoque es simple de implementar pero difícil de de escalar a medida que la aplicación crece.
```
/project-root
│
├── cmd                    # Entrypoints de la aplicación
│   └── main.go            # Archivo principal donde se inicia el servidor
│
├── pkg                    # Paquetes de funcionalidad principal
│   ├── auth               # Módulo de autenticación
│   ├── tasks              # Módulo de gestión de tareas
│   └── notifications      # Módulo de notificaciones
│
├── internal               # Código interno solo accesible dentro del módulo
│   ├── database           # Configuración de base de datos
│   └── config             # Configuración de la aplicación
│
└── go.mod                 # Archivo de dependencias de Go

```
## Basada en Microservicios
Separa la aplicación en múltiples servicios independientes, cada uno con su propia lógica de negocio y base de datos, si es necesario. Permite una escalabilidad horizontal y una implementación independiente de cada servicio, aunque puede ser más complejo de gestionar.
``` 
/project-root
│
├── user-service           # Servicio de usuarios
│   ├── cmd
│   │   └── main.go
│   ├── internal
│   │   ├── handlers       # Controladores HTTP
│   │   ├── services       # Lógica de negocio
│   │   └── database       # Conexiones a bases de datos
│   └── go.mod
│
├── product-service        # Servicio de productos
│   ├── cmd
│   ├── internal
│   ├── go.mod
│
└── order-service          # Servicio de pedidos
    ├── cmd
    ├── internal
    └── go.mod
```
## Hexagonal
Busca crear una separación estricta entre la lógica de negocio y las implementaciones externas(bases de datos, redes, UI). La idea es que la lógica de negocio principal no dependa de frameworks o detalles externos, facilitando el testing y el mantenimiento.
``` 
/project-root
│
├── cmd
│   └── main.go
│
├── internal
│   ├── domain              # Entidades del dominio
│   │   └── book.go
│   │
│   ├── application         # Lógica de negocio
│   │   └── book_service.go
│   │
│   ├── ports               # Interfaces (puertos)
│   │   └── repository.go
│   │
│   ├── adapters            # Implementaciones (adaptadores)
│   │   ├── postgres        # Base de datos PostgreSQL
│   │   └── http            # Controladores HTTP
│   │
└── go.mod
```
## Clean Architecture
Promueve la separación de responsabilidades en círculos concéntricos donde el código más interno depende del dominio y el más externo de la infraestructura. Las capas principales incluyen: entidades(dominio), casos de uso, interfaces, adaptadores y controladores externos.
``` 
/project-root
│
├── cmd
│   └── main.go
│
├── internal
│   ├── entities           # Entidades del dominio
│   │   └── user.go
│   │
│   ├── usecases           # Casos de uso
│   │   └── create_user.go
│   │
│   ├── interfaces         # Interfaces de entrada
│   │   └── http
│   │       └── handlers
│   │           └── user_handler.go
│   │
│   ├── infrastructure     # Adaptadores y controladores externos
│   │   ├── db             # Conexiones a bases de datos
│   │   └── http           # Controladores HTTP
│   │
└── go.mod
```
## Basada en Eventos
Los componentes de la aplicación se comunican entre sí mediante eventos, que se pueden publicar y suscribir en un bus de eventos. Es útil en aplicaciones de alta escalabilidad, ya que cada componente solo reacciona a eventos que le interesan.
``` 
/project-root
│
├── cmd
│   └── main.go
│
├── internal
│   ├── events             # Definición de eventos
│   │   └── order_created.go
│   │
│   ├── producers          # Publicadores de eventos
│   │   └── order_publisher.go
│   │
│   ├── consumers          # Suscriptores de eventos
│   │   └── inventory_consumer.go
│   │
│   ├── services           # Lógica de negocio
│   │   └── order_service.go
│   │
└── go.mod
```
## Arquitectura Serveless
Los desarrolladores crean funciones independientes que se ejecutan en una infraestructura gestionada, como GCP Cloud Functions. No hay necesidad de gestionar servidores y la escalabilidad se maneja automáticamente.
``` 
/project-root
│
├── functions              # Funciones individuales
│   ├── validate_image
│   │   ├── handler.go     # Lógica de la función
│   │   └── go.mod
│   │
│   ├── resize_image
│   │   ├── handler.go
│   │   └── go.mod
│   │
│   └── store_image
│       ├── handler.go
│       └── go.mod
│
└── serverless.yml         # Configuración de funciones en servicios en la nube
```