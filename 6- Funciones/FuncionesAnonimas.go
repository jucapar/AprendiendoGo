package main

import "fmt"

func main(){
    
    anonima := func(){
        fmt.Println("Me imprimo desde una funcion anonima");
    }
    anonima();
}