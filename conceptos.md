# Conceptos Fundamentales
## Tipos de Datos
**Structs**: Son tipos de datos compuestos que permiten agrupar variables bajo un ismo nombre, similar a una clase sin métodos. Los structs no soportan herencia, pero puede usar composición.

**Interfaces**: Definen un conjuto de métodos sin implementarlos, permiten un comportamiento polimórfico. 

**Métodos**: Similar a los métodos de instancia en otros lenguajes, se puede definir métodos asociados a tipos(incluidos los structs).**
## Fundamentos
**Puntero**: Permite referenciar el valor de otra variable mediante su dirección en memoria, útiles para modificar estructuras de datos grandes sin copiarlas.

**Funciones anónimas**: Son aquellas que no tienen un nombre y se pueden definir en el momento en que se usan.

## Estructuras de Datos

**Closures**: Es una función que captura variables de su entorno y las retiene aunque el entorno original termine.

**Arrays**: Colección de elementos de un solo tipo, con tamaño fijo.

**Slices**: Es una referencia a una porción de un array, tiene un tamaño dinámico.

**Maps**: Es una colección de pares clave-valor.

## Sync Package
Proporciona primitvas para manejar concurrencia.
* **Mutex**: Evita condiciones de carrera (race conditions) las cuales son un tipo de vulnerabilidad que se produce cuando varias go-rutinas acceden a recursos compartidos al mismo tiempo, sin una sincronización adecuada.
* **WaitGroup**: Permite esperar a que un grupo de goroutines terminen.
* **Once**: Garantiza que una función se ejecute solo una vez.

## Manejo de Errores
Go utiliza el manejo de errores explícito, se pueden crear errores personalizados usando **erros.New()** o **fmt.Errorf()**.

## Testing y Mocks
Go tiene un paquete de testing integrado para escribir pruebas(testing).
```go
func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5 but got %d", result)
    }
}
```
Así mismo existe la biblioteca **GoMock** para crear mocks de interfaces en pruebas unitarias.
```go
ctrl := gomock.NewController(t)
defer ctrl.Finish()
mockService := NewMockMyService(ctrl)
mockService.EXPECT().DoSomething().Return(nil)
```