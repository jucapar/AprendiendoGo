package main

import "fmt"

func main() {
	//Declaracion de un array  nombreArray [tamaño]TipoDeDato
	var a [2]string
	a[0] = "Hola"
	a[1] = "Mundo"
	fmt.Println(a[0] + " " + a[1])

	//Declaracion de un array inicializado
	diasSemana := [7]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}
	fmt.Println(diasSemana[4])
}
