package main

import "fmt"

func main() {
    fmt.Println("Informe a temperatura em ºF: ")
    var grausFahrenheit float64
    fmt.Scanf("%f", &grausFahrenheit)

    grausCelsius := (grausFahrenheit - 32) * 5/9

    fmt.Println("Temperatura em ºC")
    fmt.Println(grausCelsius)
}
