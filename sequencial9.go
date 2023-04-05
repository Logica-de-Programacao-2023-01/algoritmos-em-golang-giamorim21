package main

//Faça um algoritmo que leia o peso de uma pessoa em quilos 
//e converta para libras.

import "fmt"

func main() {
     var peso float64
     fmt.Println("Digite seu peso (em quilos):")
     fmt.Scan(&peso)
     
     
     fmt.Println("O seu peso em libras eh :",peso*2.20462)
}
