package main

//Faça um algoritmo que leia a altura e o sexo de uma pessoa 
//e mostre se ela está abaixo, dentro ou acima do peso ideal, 
//utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais

import "fmt"

func main() {
  var altura, peso float64
  var sexo string
  fmt.Println("Digite sua altura:")
  fmt.Scan(&altura)
  
  fmt.Println("Digite seu sexo (M para masculino , F para feminino")
  fmt.Scan(&sexo)
  
  if sexo == "M"{
      fmt.Println("peso ideal = ",(0.75* altura) - 62.5)
  } else if sexo == "F"{
      fmt.Println("peso ideal =",(0.67 * altura) - 56)
  }
  
  fmt.Println("Digite seu peso:")
  fmt.Scan(&peso)
  
  imc := peso /(altura*altura)
  
  if imc < 18.5 {
      fmt.Println("Você está abaixo do peso")
  } else if imc > 25 {
      fmt.Println("Você está acima do peso")
  } else {
      fmt.Println("Você está dentro do peso")
  }}

