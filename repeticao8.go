package main
//Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.

import "fmt"

func main (){
    var x int 
    fmt.Println("Digite um número:")
    fmt.Scan(&x)
    
    fmt.Println("os divisores de x são:", x)
    
    for i := 1; i <= x; i++{
        if x % i == 0{
            fmt.Println(i)
        }
    }
}
