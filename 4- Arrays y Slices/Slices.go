package main

import "fmt"


func main(){
    
    //Declaracion de un Slice make(tipoDeDato,longitud,capacidad)
    //Longitud - Capacidad del slice inicial
    //Capacidad - Capacidad maxima del slice, al superarse se añade espacio de memoria igual a la capacidad inicial
    slice := make([]string, 2, 5)
    slice[0] = "Hola";
    slice[1] = "Mundo";
    
    fmt.Println(slice[0]);
    fmt.Println(slice[1]);
}