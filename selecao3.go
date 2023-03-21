package main

import "fmt"

// Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.
func main() {
	var x int
	fmt.Println("Digite o numero x")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println(x, "é par")
	} else if x%2 != 0 {
		fmt.Println(x, "é impar")
	}

}
