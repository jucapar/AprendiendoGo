package main

import "fmt";

func main(){
   
    Saludar("Juan Carlos","Lucia","Cristina");
}

func Saludar(nombres ...string){
    for _,valor := range nombres{
        fmt.Println("Hola", valor);
    }
}