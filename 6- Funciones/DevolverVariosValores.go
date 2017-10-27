package main

import "fmt";

func main(){
    numeros := []int8{53,21,12,99,33};
    max,min := MaxYMin(numeros);
    fmt.Println("Maximo:",max);
    fmt.Println("Minimo:",min);
}

func MaxYMin(numeros []int8) (max int8,min int8){
   
    max = numeros[0];
    min = numeros[0];
    for _,valor := range numeros{
        if (valor > max){
            max = valor;
        }
        
        if(valor < min){
            min = valor;
        }
        
    }
    
    return 
}