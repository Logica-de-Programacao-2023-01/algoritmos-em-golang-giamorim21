package main

//Faça um algoritmo que leia o número de dias trabalhados por um funcionário
//e o valor da sua diária e calcule o seu salário.

import "fmt"

func main() {
     var dias, diaria int
     fmt.Println("Digite o numero de dias trabalhados (em dias) :")
     fmt.Scan(&dias)
     
     fmt.Println("Digite o valor da sua diaria (em reais):")
     fmt.Scan(&diaria)
     
     novoSalario := dias*diaria
     
     fmt.Println("O valor do novo salario do funcionario eh =", novoSalario ,"reais")
}
