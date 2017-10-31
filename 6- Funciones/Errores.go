package main

import (
        "fmt"
        "errors" //Importamos la libreria errors
       )
func main(){
    
    resultado,err := division(1000,0);
    
    if(err != nil){
        fmt.Println("Error:",err);
        return;
    }
        fmt.Println(resultado);
   
    
}

func division(dividendo float64, divisor float64) (resultado float64, err error){
    
    if(divisor == 0){
        err = errors.New("No se puede dividir entre 0");
    }else{
        resultado = dividendo / divisor;
    }
    return
}