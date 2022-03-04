package main

import "fmt"

func main() {

	var nombre1 string

	nombre1 = "Hola Mundo persona"
	nombre := 26.5

	fmt.Println(nombre1)
	fmt.Println(nombre)

	nombre = 24
	fmt.Println(nombre)

	result := sumar(10, 20)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

AAAAAA
