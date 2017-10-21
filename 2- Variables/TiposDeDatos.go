package main

import "fmt"
/*
 TIPOS ENTEROS
 
1 	uint8
Números enteros sin signo de 8 bits (0 a 255)
2 	uint16
Enteros sin signo de 16 bits (0 a 65535)
3 	uint32
Enteros sin signo de 32 bits (0 a 4294967295)
4 	uint64
Enteros sin signo de 64 bits (0 a 18446744073709551615)
5 	INT8
Enteros con signo de 8 bits (-128 a 127)
6 	Int16
Enteros con signo de 16 bits (-32768 a 32767)
7 	int32
Enteros con signo de 32 bits (-2147483648 a 2147483647)
8 	Int64
Enteros con signo de 64 bits (-9223372036854775808 a 9223372036854775807) 

TIPOS FLOAT

1 	float32
IEEE-754 números de coma flotante de 32 bits
2 	float64
IEEE-754 números de coma flotante de 64 bits
3 	complex64
Los números complejos con Float32 partes real e imaginaria
4 	complex128
Los números complejos con float64 partes real e imaginaria 


OTROS TIPOS NUMERICOS

1 	byte
mismo que uint8
2 	runa
mismo que int32
3 	uint
32 o 64 bits
4 	int
mismo tamaño que uint
5 	UIntPtr
un entero sin signo para almacenar los bits no interpretados de un valor de puntero 
*/
func main(){
    
    /*Tipos basicos*/
    var entero int = 4;
    var real float32 = 4.5;
    var cadena string = "Soy una cadena";
    var booleano bool = true;
    
    fmt.Println(entero);
    fmt.Println(real);
    fmt.Println(cadena);
    fmt.Println(booleano);
}