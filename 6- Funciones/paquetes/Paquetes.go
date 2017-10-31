package main

import (
    "fmt"
    "./despedida"
    "./saludo"
)

func main(){
    
    saludo.Saludar("Don Pepito");
    saludo.Saludar("Don Jose");
    fmt.Println("Pasó usted por mi casa");
    fmt.Println("Por su casa yo pasé");
    fmt.Println("Vió usted a mi abuela");
    fmt.Println("A su abuela yo la vi");
    despedida.Despedirse("Don Pepito");
    despedida.Despedirse("Don Jose");
}