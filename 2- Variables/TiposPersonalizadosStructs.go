package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	altura float32
}

func main() {

	var yo = Persona{
		nombre: "Jucapar",
		edad:   28,
		altura: 1.72}

	fmt.Println(yo)

	fmt.Print("Nombre:")
	fmt.Println(yo.nombre)
	fmt.Print("Edad:")
	fmt.Println(yo.edad)
	fmt.Print("Altura:")
	fmt.Println(yo.altura)
}
