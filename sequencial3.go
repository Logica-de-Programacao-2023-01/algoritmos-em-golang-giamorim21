package main

func main() {
 var peso float64
    var altura float64
    fmt.Println("Digite seu peso ( em kg ): ")
    fmt.Scan(&peso)
    
    fmt.Println("Digite sua altura ( em metros ): ")
    fmt.Scan(&altura)
    
    fmt.Println("peso/(altura*altura)", peso/(altura*altura))
    fmt.Println (" O seu imc eh de :",peso/(altura*altura) )
}
