package main

import "fmt"

func main() {

	dia := 2

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("No es un dia de la semana")
	}

	switch dia {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estas entre semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un dia de la semana")
	}
}
