package main

import "fmt"

func main() {

	var nombre1 string

	nombre1 = "Hola Mundo persona 2444433336"
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


func sumar2(a, b int) int {
	return a + b
}
