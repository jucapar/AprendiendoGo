package main

import "fmt"

func main(){
    a := 3;
    fmt.Println("Valor antes de duplicar:",a);
    duplicar(&a);
    fmt.Println("Valor despues de duplicar:",a);
    
}
func duplicar(a *int){
    *a *= 2;
}