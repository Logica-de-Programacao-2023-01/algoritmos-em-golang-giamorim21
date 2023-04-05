package main
//Faça um algoritmo que leia a idade de uma pessoa 
//e mostre a sua classificação, de acordo com a seguinte tabela:
//até 9 anos: mirim
//10 a 13 anos: infantil
//14 a 17 anos: juvenil
//maiores de 18 anos: adulto

import "fmt"

func main (){
    var idade int 
    fmt.Println("Digite a sua idade:")
    fmt.Scan(&idade)
    
    switch {
    case idade <= 9:
        fmt.Println("Mirim")
    case idade <= 13:
        fmt.Println("Infantil")
    case idade <= 17:
        fmt.Println("Juvenil")
    default:
        fmt.Println("Adulto")
    }
    
}
