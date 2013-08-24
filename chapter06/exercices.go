package main

import "fmt"

func main() {
    x := make([]int, 3, 9)
    fmt.Println(len(x))

    y := [6]string{"a", "b", "c", "d", "e", "f"}
    fmt.Println(y[2:5])

    z := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }

    var menor int = z[0]
    var maior int = z[0]

    for _, value := range z {
        if menor > value {
            menor = value
        }
    }
    for _, value := range z {
        if maior < value {
            maior = value
        }
    }

    fmt.Println(maior)
    fmt.Println(menor)
}
