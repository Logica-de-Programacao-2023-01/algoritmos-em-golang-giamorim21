package main
//Faça um algoritmo que leia dois números inteiros 
//e mostre o resultado da multiplicação entre eles,
//se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.

import "fmt"

func main() {
    var x, y int
    fmt.Println("Digite um numero para x:")
    fmt.Scan(&x)
    
    fmt.Println("Digite um numero para y:")
    fmt.Scan(&y)
    
    if x > 0 && y > 0{
        fmt.Println("os dois são positivos --> x*y = ", x*y)
    } else if x < 0 || y < 0 {
        fmt.Println("pelo menos um deles é negativo --> x+y = ",x+y)
    }
}
