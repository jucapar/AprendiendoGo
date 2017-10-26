package main

import "fmt"


func main(){
    
    //Declaracion de un Slice make(tipoDeDato,longitud,capacidad)
    //tipoDeDato - Apuntador a un array
    //Tamaño - Capacidad del slice inicial
    //Capacidad(Opcional) - Capacidad maxima del slice, al superarse se duplica la capacidad del slice
    slice := make([]string, 0)
    slice = append(slice,"Hola","Mundo");
    
    fmt.Println(slice);
    fmt.Printf("El tamaño del slice es de %d y su capacidad es de %d",len(slice),cap(slice));
}