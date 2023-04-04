package main

//Faça um algoritmo que leia três números inteiros
//e mostre a soma entre eles.

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Println("Digite um numero para x:")
	fmt.Scan(&x)
	fmt.Println("Digite um numero para y:")
	fmt.Scan(&y)
	fmt.Println("Digite um numero para z:")
	fmt.Scan(&z)

	fmt.Println(" x + y + z =", x+y+z)
}