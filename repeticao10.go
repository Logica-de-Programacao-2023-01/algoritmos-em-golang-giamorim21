package main
//Faça um algoritmo que leia vários números inteiros e mostre o maior deles. 
//A leitura deve ser interrompida quando for lido o valor zero.

import "fmt"

func main (){
   var x, maior int
    for {
        fmt.Println("Digite um número inteiro (0 para sair): ")
        fmt.Scan(&x)
          
          if x == 0{
              break
          }
          
          if x > maior {
              maior = x
          }
    }
    
    fmt.Println("O maior número digitado foi :", maior)
}
