package main

//Faça um algoritmo que leia um número inteiro
//e mostre o seu sucessor e antecessor.

import "fmt"

func main() {
    var x int

    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&x)

    antecessor := x - 1
    sucessor := x + 1

    fmt.Printfn("O antecessor de %d é %d e o sucessor é %d", x , antecessor, sucessor)
}
