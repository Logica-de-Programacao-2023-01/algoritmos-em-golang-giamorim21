package main
//Faça um algoritmo que leia vários números inteiros
//e mostre a média aritmética entre eles.
//A leitura deve ser interrompida quando for lido o valor zero.

import "fmt"

func main (){
    var soma, quantidade, x int
    
    fmt.Println("Insira um número: ")
    fmt.Scan(&x)
    
    if x != 0 {
        soma += x
        quantidade++
    }
    
    for x != 0 {
        fmt.Println("Insira um número: ")
        fmt.Scan(&x)
        
        if x != 0 {
            soma += x 
            quantidade++
        }
    }
    
    média := soma / quantidade 
    fmt.Println("A média eh", média)
}
