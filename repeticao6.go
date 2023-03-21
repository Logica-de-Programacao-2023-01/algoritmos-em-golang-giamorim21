package main

import "fmt"

// Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.
func main() {
	var x int
	fmt.Println("Digite o numero x ")
	fmt.Scan(&x)

	for i := 1; i <= 10; i++ {
		fmt.Println(x, "x", i, "=", x*i)

	}

}
