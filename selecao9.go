package main
//Faça um algoritmo que leia três números reais e mostre-os em ordem crescente

import (
    "fmt"
    "sort"
)

func main() {
    var num1, num2, num3 float64

    fmt.Println("Digite o primeiro número:")
    fmt.Scanln(&num1)

    fmt.Println("Digite o segundo número:")
    fmt.Scanln(&num2)

    fmt.Println("Digite o terceiro número:")
    fmt.Scanln(&num3)

    numeros := []float64{num1, num2, num3}
    sort.Float64s(numeros)

    fmt.Println("Os números em ordem crescente são:", numeros)
}
