package main
//Faça um algoritmo que leia o salário de um funcionário
//e calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00;
//ou de 5% se o salário for maior ou igual a R$ 1000,00.

import "fmt"

func main() {
    var salário float64
    fmt.Println("Digite o salário do funcionário:")
    fmt.Scan(&salário)
    
    if salário < 1000 {
        fmt.Println("Com o aumento o salário agora eh de: ",(salário*0.1) + salário )
    } else if salário > 1000 {
        fmt.Println("Com o aumento o salário agora eh  de:",(salário*0.05) + salário )
    }
}
