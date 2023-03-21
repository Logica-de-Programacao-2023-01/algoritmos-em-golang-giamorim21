package main

import "fmt"

//Faça um algoritmo que imprima os números ímpares de 1 a 19.

func main() {
	for i := 1; i <= 19; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}

	}

}
