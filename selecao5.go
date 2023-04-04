package main
//Faça um algoritmo que leia um número inteiro 
//e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

import "fmt"

func main() {
    var x int 
    fmt.Println("Digite um número:")
    fmt.Scan(&x)
    
    if x % 3 == 0 && x % 5 == 0 {
        fmt.Println("Esse número eh múltiplo de 3 e de 5 simultaneamente")
    } else {
        fmt.Println ("Esse número não eh múltiplo de 3 e de 5 simultaneamente")
    }
}
