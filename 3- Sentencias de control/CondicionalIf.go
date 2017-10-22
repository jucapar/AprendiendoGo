package main

import "fmt"

func main() {

	edad := 15

	if edad < 14 {
		fmt.Println("Es un niño")
	} else if edad < 18 {
		fmt.Println("Es un adolescente")
	} else {
		fmt.Println("Es un adulto")
	}

}
