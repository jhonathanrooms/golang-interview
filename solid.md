# S.O.L.I.D
Son cinco principios fundamentales para el diseño de software orientado a objetos. Promueven la creación de código modular, fácil de mantener y escalable. 
## Resumen
* S **SRP**: Una clase o función debe tener una única responsabilidad.
* O **OCP**: El código debe ser extensible sin modificar el existente.
* L **LSP**: Las subclases deben ser sustituibles sin afectar la funcionalidad.
* I **ISP**: Divide las interfaces grandes en interfaces específicas.
* D **DIP**: Las dependencias deben estar basadas en abstracciones, no en implementaciones.
En Go, aunque la orientación a objetos se aplica de una manera menos explícita que en otros lenguajes, los principios SOLID aún se pueden seguir. Aplicarlos mejora la mantenibilidad, flexibilidad y claridad del código.
## S - Single Responsibility Principiple
Una clase(o componente) debe tener una única responsabilidad o motivo para cambiar, debe encargarse de hacer una sola cosa.
**Uso en Go**: Las funciones, estructuras y métodos deben tener un propósito claro y cumplir una única función.
``` go
// Violación del SRP: La misma estructura tiene responsabilidades de almacenamiento y validación.
type User struct {
    Name  string
    Email string
}

func (u *User) SaveToDatabase() {
    // Código para guardar al usuario en la base de datos
}

func (u *User) ValidateEmail() bool {
    // Código para validar el correo electrónico
    return true
}

// Aplicación del SRP: Separar responsabilidades en diferentes componentes.
type User struct {
    Name  string
    Email string
}

type UserRepository struct{}

func (r *UserRepository) Save(user User) {
    // Código para guardar el usuario en la base de datos
}

type UserValidator struct{}

func (v *UserValidator) ValidateEmail(email string) bool {
    // Código para validar el correo electrónico
    return true
}
```
2. O - Open/Closed Principle
Los componentes de software deben estar abiertos para extensión, pero cerrados para modificación. Debe ser posibke agregar nuevas funcionalidades sin modificar el código existente.
**Uso en Go**: Se permite aplicar el principio mediante el uso de interfaces. Se pueden crear nuevas implementaciones de una interfaz sin cambiar el código que usa dicha interfaz.
``` go
// Interface para diferentes tipos de notificaciones.
type Notifier interface {
    Notify(message string)
}

// Implementación de Notifier para correo electrónico.
type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) {
    // Código para enviar un correo electrónico
}

// Implementación de Notifier para SMS.
type SMSNotifier struct{}

func (s SMSNotifier) Notify(message string) {
    // Código para enviar un SMS
}

func SendAlert(notifier Notifier, message string) {
    notifier.Notify(message)
}

// Agregar nuevas formas de notificación, como Slack, solo requiere
// implementar la interfaz Notifier sin modificar SendAlert.
```
3. L - Liskov Substitution Principle
Los objetos de una clase derivada deben poder reemplazar a los objetos de la clase base sin alterar el funcionamiento del programa.
**Uso en Go**: Cuando una estructura satisface una interfaz, debería poder usarse en lugar de cualquier otra estructura que también satisfaga esa interfaz sin problemas de compatibilidad.
``` go
// Interface para cálculo de áreas.
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Se pueden usar tanto Rectangle como Circle donde Shape sea esperado.
func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

// Ambos cumplen LSP, ya que ambas estructuras pueden ser usadas con la interfaz Shape.
```
4. I - Interface Segregation Principle.
Los clientes no deberían estar forzados a depender de interfaces que no usan. En vez de tener una única interfaz con muchos métodos, es mejor dividirla en interfaces más pequeñas y específicas.
**Usos en Go**: Se aplican interfaces pequeñas y específicas en lugar de una interfaz grande con métodos innecesarios.
``` go
// Violación del ISP: Una interfaz con demasiados métodos.
type Worker interface {
    Work()
    Sleep()
    Eat()
}

type Human struct{}

func (h Human) Work() { /* Implementación de trabajo */ }
func (h Human) Sleep() { /* Implementación de sueño */ }
func (h Human) Eat() { /* Implementación de comida */ }

// Aplicación del ISP: Interfaces más pequeñas y específicas.
type Worker interface {
    Work()
}

type Sleeper interface {
    Sleep()
}

type Eater interface {
    Eat()
}

type Robot struct{}

func (r Robot) Work() { /* Implementación de trabajo */ }

type Human struct{}

func (h Human) Work() { /* Implementación de trabajo */ }
func (h Human) Sleep() { /* Implementación de sueño */ }
func (h Human) Eat() { /* Implementación de comida */ }
```
5. D - Dependency Inversion Principle
Las clases de alto nivel no deberían depender de las clases de bajo nivel. Ambas deben depender de abstracciones(interfaces). Las abstracciones no deben depender de los detalles, los detalles deben depender de las abstracciones.
**Usos en Go**: Las dependencias se pasan como interfaces en lugar de estructuras concretas, permitiendo que los componentes de alto nivel dependan de las interfaces.
``` go
// Violación del DIP: HighLevel depende directamente de LowLevel.
type FileLogger struct{}

func (l FileLogger) Log(message string) {
    // Código para registrar mensaje en archivo
}

type App struct {
    logger FileLogger
}

func (a App) DoSomething() {
    a.logger.Log("Doing something...")
}

// Aplicación del DIP: HighLevel depende de una abstracción (Logger).
type Logger interface {
    Log(message string)
}

type FileLogger struct{}

func (l FileLogger) Log(message string) {
    // Código para registrar mensaje en archivo
}

type App struct {
    logger Logger
}

func (a App) DoSomething() {
    a.logger.Log("Doing something...")
}

// Ahora es fácil cambiar FileLogger por otro tipo de Logger, como ConsoleLogger.
```