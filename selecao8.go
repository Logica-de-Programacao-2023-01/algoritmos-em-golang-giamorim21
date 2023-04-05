package main
//Faça um algoritmo que leia um número inteiro entre 1 e 7 
//e mostre o dia da semana correspondente (1 = domingo, 2 = segunda-feira, etc.).

import "fmt"

func main() {
   var x int
   fmt.Println("Digite um número entre 1 e 7")
   fmt.Scan(&x)
   
   switch x {
       case 1:
         fmt.Println("Hoje é domingo")
       case 2:
         fmt.Println("Hoje é segunda")
       case 3:
         fmt.Println("Hoje é terça")
       case 4:
         fmt.Println("Hoje é quarta")
       case 5:
         fmt.Println("Hoje é quinta")
       case 6:
         fmt.Println("Hoje é sexta")
       case 7:
         fmt.Println("Hpje é sábado")
   }
}
