package main

//Faça um algoritmo que leia um número inteiro
//e mostre o seu dobro, triplo e quádruplo.

import "fmt"

func main() {
    var x int
    
    fmt.Println("Digite um numero para x:")
    fmt.Scan(&x)
    
    fmt.Println("x*2 =", x*2)
    fmt.Println("x*3 =", x*3)
    fmt.Println("x*4 =", x*4)
}

