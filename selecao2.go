package main

import "fmt"

// Faça um algoritmo que leia três números inteiros e mostre o menor deles.
func main() {
	var x int
	var y int
	var z int
	fmt.Println("Digite o numero x")
	fmt.Scan(&x)

	fmt.Println("Digite o numero y")
	fmt.Scan(&y)

	fmt.Println("Digite o numero z")
	fmt.Scan(&z)
	if x < y && x < z {
		fmt.Println(x, "este é o menor numero ")
	} else if y < x && y < z {
		fmt.Println(y, "este é o menor numero")
	} else if z < x && z < y {
		fmt.Println(z, "este é o menor numero")
	}
}
