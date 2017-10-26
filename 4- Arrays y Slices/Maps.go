package main

import "fmt"


func main(){
    
    //Mapas: Similares a los array asociativos de php
    //Declaracion nombreVariable := make(map[tipoDatoClave]tipoDatoValor)
    mapa := make(map[string]string);
    mapa["es"] = "español";
    mapa["en"] = "ingles";
   
    //Otra forma de declaracion
    mapa2 := map[string]string{
        "ch":"chino",
        "fr":"frances",
        };
        
    fmt.Println(mapa);
    fmt.Println(mapa2["ch"]);
    //Borrar una posicion del mapa
    delete(mapa2,"ch");
    fmt.Println(mapa2);
    //Si intentamos imprimir una posicion que no existe nos imprimirá vacio
    fmt.Println(mapa["jp"]);
    //No podremo saber si el indice existe en el map y el valor está vacio, por
    //lo que tendremos que comprobarlo con un if
    
    if idioma,ok := mapa["jp"]; ok {
        fmt.Println("La posicion existe y vale", idioma);
    }else{
        fmt.Println("La posicion no existe");
    }
    //el mapa nos devuelve el valor de la posicion (que se guardará en idioma) y true si la
    //posicion existe (que se guardará en ok), luego evaluamos el ok para la comprobacion
}