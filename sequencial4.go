package main

// Faça um algoritmo que leia três números reais 
//e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.

import "fmt"

func main() {
[14:16, 31/03/2023] giovana amorim: var num1, num2, num3 float64

    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&num1)
    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&num2)
    fmt.Print("Digite o terceiro número: ")
    fmt.Scanln(&num3)

    fmt.Println("(num1*2 + num2*3 + num3*5)/10",(num1*2 + num2*3 + num3*5)/10)
    
    fmt.Println("A média ponderada é: ",(num1*2 + num2*3 + num3*5)/10 )
}
