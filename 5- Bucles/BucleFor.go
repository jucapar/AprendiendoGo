package main

import "fmt"

func main(){
    
    for i := 0;i < 5;i++{
        fmt.Println(i);
    }
    
    matriz := [3][3]int{};
    valor := 0;
    for i := 0;i< len(matriz); i++{
        for j := 0;j< len(matriz[0]); j++{
            valor++;
            matriz[i][j] = valor;
        }
    }
    
     for i := 0;i< len(matriz); i++{
        for j := 0;j< len(matriz[0]); j++{
            
            fmt.Print(matriz[i][j]," ");
            
        }
        fmt.Println();
    }
    
    a := 0;
    for a < 10{
        fmt.Print(a," ");
        a++;
    }
    fmt.Println();
    
    /*
    Bucle infinito
    for{
        fmt.Println("Hola");
    }
    */
    
    //for range
    
    nombres := []string{"Juan Carlos","Lucia","Cristina"};
    
    for i,valor := range nombres{
        //El for devuelve el indice del array que se almacenará en i
        // y el valor que se almacenará en valor
        fmt.Println("Posicion ",i,":",valor);
    }
    
    
    //Si no necesitamos utilizar el indice
    for _,valor := range nombres{
        fmt.Println(valor);
    }
    
    //Podemos recorrer cadenas caracter a caracter mediante el bucle for
    frase := "Hola mundo";
    for i,valor := range frase{
        fmt.Println("Letra",i,":",string(valor));
        //Si no casteamos el valor, nos devuelve su numero ascii correspondiente al caracter
    }
}