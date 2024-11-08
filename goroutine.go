package main

import (
	"fmt"
)

func myFunction() {
	fmt.Println("Hello from goroutine!")
}

func goroutine() {
	go myFunction() // Lanza una nueva goroutine
	fmt.Println("Hello from main!")
}
