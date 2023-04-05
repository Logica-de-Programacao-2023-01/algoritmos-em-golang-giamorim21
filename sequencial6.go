package main

//Faça um algoritmo que leia o salário de um funcionário 
//e calcule o seu novo salário com um aumento de 15%.

import "fmt"

func main() {
   var salario float64
   fmt.Println("Digite qual eh o seu salario antigo:")
   fmt.Scan(&salario)
   
   novoSalario := (salario * 0.15) + salario
   
   fmt.Println("O novo salário do funcionário é ", novoSalario)
    
}
