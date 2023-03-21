package main

import "fmt"

//Faça um algoritmo que leia dois números inteiros e mostre o maior deles.

func main() {
	var x int
	var y int
	fmt.Println("Digite o numero x")
	fmt.Scan(&x)

	fmt.Println("Digite o numero y")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("x é maior que y")
	} else {
		fmt.Println("y é maior que x")

	}

}
