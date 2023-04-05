package main

// Faça um algoritmo que leia a idade de uma pessoa em anos 
//e mostre a idade em dias.   

import "fmt"

func main() {
   var idade int
   fmt.Println("Digite a sua idade (em anos):")
   fmt.Scan(&idade)
   
   fmt.Println("Sua idade em dias =", idade*365)
}
