package main

import "fmt"

func main() {
    fmt.Println("Digite um valor em pés:")
    var ft float64
    fmt.Scanf("%f", &ft)

    m := ft * 0.3048

    fmt.Println("Valor em metros")
    fmt.Println(m)
}
