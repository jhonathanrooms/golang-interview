# Programación Orientada a Objetos

En Go la POO es un poco diferente a lenguajes como Python o Java, ya que Go no tiene clases ni herencia. Sin embargo se puede implementar conceptos de POO utilizando structs, interfaces y métodos(Una función que tiene un argumento receptor especial).
## Resumen de POO en Go
* **Encapsulamiento**: Control de acceso con nombres en mayúscula o minúscula.
* **Abstracción**: Definición de interfaces que especifican qué métodos deben implementarse.
* **Herencia (Composición)**: Inclusión de estructuras dentro de otras para reutilizar métodos.
* **Polimorfismo**: Uso de interfaces para trabajar con diferentes tipos que implementan los mismos métodos.
## Conceptos
### Encapsulamiento
Se maneja mediante el uso de letras mayúsculas y minúsculas en los nombres de los campos y métodos. En Go, los nombres que empiezan con mayúscula son públicos, mientras que los nombres en minúscula son privados.
```go 
package main

import "fmt"

// Definimos una estructura Persona con un campo privado y uno público
type Persona struct {
	nombre string // campo privado
	Edad   int    // campo público
}

// Método para establecer el nombre (setter)
func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

// Método para obtener el nombre (getter)
func (p *Persona) GetNombre() string {
	return p.nombre
}

func main() {
	p := Persona{Edad: 30}
	p.SetNombre("Juan")
	fmt.Println("Nombre:", p.GetNombre()) // Accedemos al nombre de forma encapsulada
	fmt.Println("Edad:", p.Edad)          // Campo público, se accede directamente
}
```
### Abstracción
Se logra utilizando interfaces, una interfaz define métodos que una estructura debe implementar, sin dar detalles de implementación.
```go 
package main

import "fmt"

// Definimos una interfaz Animal con un método Speak
type Animal interface {
	Speak() string
}

// Estructuras que implementan la interfaz Animal
type Perro struct{}
type Gato struct{}

func (p Perro) Speak() string {
	return "Guau"
}

func (g Gato) Speak() string {
	return "Miau"
}

func comunicar(animal Animal) {
	fmt.Println(animal.Speak())
}

func main() {
	perro := Perro{}
	gato := Gato{}

	comunicar(perro) // Output: Guau
	comunicar(gato)  // Output: Miau
}
```
### Herencia
No hay herencia de clases como tal, pero se puede simular la herencia utilizando composición. Esto significa que una estructura puede incluir otra como campo, permitiendo reutilizar el comportamiento de esta última.
```go 
package main

import "fmt"

// Estructura Común con comportamiento básico
type Vehiculo struct {
	Marca string
}

func (v Vehiculo) Encender() {
	fmt.Println("Vehículo encendido")
}

// Carro "hereda" de Vehiculo mediante composición
type Carro struct {
	Vehiculo // Campo embebido
	Modelo   string
}

func main() {
	c := Carro{
		Vehiculo: Vehiculo{Marca: "Toyota"},
		Modelo:   "Corolla",
	}

	fmt.Println("Marca:", c.Marca)
	fmt.Println("Modelo:", c.Modelo)
	c.Encender() // Utiliza el método del campo embebido
}
```
### Polimorfismo
Se logra utilizando interfaces, si varias estructuras implementan los mismos métodos de una interfaz pueden usarse de manera intercambiable.
```go 
package main

import "fmt"

// Interfaz Empleado con un método CalcularSalario
type Empleado interface {
	CalcularSalario() float64
}

// Estructura para Empleados por Hora
type EmpleadoPorHora struct {
	HorasTrabajadas int
	TarifaPorHora   float64
}

func (e EmpleadoPorHora) CalcularSalario() float64 {
	return float64(e.HorasTrabajadas) * e.TarifaPorHora
}

// Estructura para Empleados Fijos
type EmpleadoFijo struct {
	SalarioMensual float64
}

func (e EmpleadoFijo) CalcularSalario() float64 {
	return e.SalarioMensual
}

func calcularSalarioMensual(e Empleado) {
	fmt.Printf("Salario mensual: %.2f\n", e.CalcularSalario())
}

func main() {
	empleado1 := EmpleadoPorHora{HorasTrabajadas: 160, TarifaPorHora: 15}
	empleado2 := EmpleadoFijo{SalarioMensual: 3000}

	calcularSalarioMensual(empleado1) // Output: Salario mensual: 2400.00
	calcularSalarioMensual(empleado2) // Output: Salario mensual: 3000.00
}
```