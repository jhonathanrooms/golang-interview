# Patrones de Diseño
Son soluciones habituales a problemas que ocurren con frecuencia en el diseño de software. Son planos prefabricados que se pueden personalizar para reslver un problema de diseño recurrente.

Todos los patrones pueden clasificarse por su propósito, siendo los tres grupos generales:
* **Creacionales**: Proporcionan mecanismos de creación de objetos que incrementan la flexibilidad y la reutilización de código existente.
* **Estructurales**: Explican cómo ensamblar objetos y clases en estructuras más grandes a la vez que se mantiene la flexibilidad y eficiencia de la estructura.
* **Comportamiento**: Se encargan de una comunicación efectiva y la asignación de responsabilidades entre objetos.
## Singleton (Creacional)
Garantiza que una clase solo tenga una instancia y proporciona un punto de acceso global a ella. En Go, se suele implementar utilizando variables y **sync.Once** para asegurarse de que se inicialice solo una vez.
``` go
package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2) // Output: true
}
```
## Factory (Creacional)
Proporciona una interfaz para crear objetos en una superclase, pero permite a las subclases alterar el tipo de objetos que se crean. Es útil para evitar dependencias directas con clases concretas y para crear objetos basados en condiciones.
``` go
package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string { return "Woof!" }
func (c Cat) Speak() string { return "Meow!" }

func AnimalFactory(animalType string) Animal {
	switch animalType {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}

func main() {
	dog := AnimalFactory("dog")
	cat := AnimalFactory("cat")

	fmt.Println(dog.Speak()) // Output: Woof!
	fmt.Println(cat.Speak()) // Output: Meow!
}
```
## Strategy (Comportamiento)
Permite definir una familia de algortimos, encapsular cada yno y hacerlos intercambiables, esto permite que el algoritmo varíe independientemente del cliente que lo utiliza.
``` go
package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCard struct{}
type PayPal struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %0.2f using Credit Card", amount)
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %0.2f using PayPal", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) string {
	return p.strategy.Pay(amount)
}

func main() {
	context := PaymentContext{}

	context.SetStrategy(&CreditCard{})
	fmt.Println(context.Pay(100.0)) // Output: Paid 100.00 using Credit Card

	context.SetStrategy(&PayPal{})
	fmt.Println(context.Pay(200.0)) // Output: Paid 200.00 using PayPal
}
```
## Observer (Comportamiento)
Define una relación de dependencia uno a muchos, de forma que cuando un objeto cambia de estado, todos sus dependencias son notificados y actualizados automáticamente.
``` go
package main

import "fmt"

type Observer interface {
	Update(string)
}

type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) SetState(state string) {
	s.state = state
	for _, observer := range s.observers {
		observer.Update(state)
	}
}

type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(state string) {
	fmt.Printf("%s received state change: %s\n", co.name, state)
}

func main() {
	subject := &Subject{}
	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.SetState("new state") // Output: Observer1 received state change: new state
	                               //         Observer2 received state change: new state
}
```
## Adapter (Estructural)
Permite que dos interfaces incompatibles trabajen juntas. Actúa como un puente entre dos interfaces diferentes y permite que una clase se ajuste a una interfaz requerida sin modificar su código original.
``` go
package main

import "fmt"

// Target es la interfaz esperada.
type Target interface {
	Request() string
}

// Adaptee es la interfaz incompatible.
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Hello from Adaptee!"
}

// Adapter hace compatible Adaptee con Target.
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}
	fmt.Println(adapter.Request()) // Output: Hello from Adaptee!
}
```
## Decorator (Estructural)
Permite añadir dinámicamente comportamientos adicionales a un objeto sin modificar su estructura. En Go, se puede implementar decorando las funciones que se envuelven en otras funciones.
``` go
package main

import "fmt"

// Component es la interfaz base.
type Notifier interface {
	Send(message string)
}

// ConcreteComponent es el objeto base que será decorado.
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("Email notification sent: %s\n", message)
}

// Decorator agrega comportamiento a Notifier.
type SMSDecorator struct {
	notifier Notifier
}

func (s *SMSDecorator) Send(message string) {
	s.notifier.Send(message)
	fmt.Printf("SMS notification sent: %s\n", message)
}

func main() {
	email := &EmailNotifier{}
	smsEmail := &SMSDecorator{notifier: email}

	email.Send("Hello World!")      // Output: Email notification sent: Hello World!
	smsEmail.Send("Hello World!")   // Output: Email notification sent: Hello World!
	                                //         SMS notification sent: Hello World!
}
```