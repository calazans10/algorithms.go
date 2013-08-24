package main

import "fmt"

func main() {
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr)

    x := 1.5
    square(&x)
    fmt.Println(x)

    a := 1
    b := 2
    swap(&a, &b)
    fmt.Println("x =", a)
    fmt.Println("y =", b)
}

func one(xPtr *int) {
    *xPtr = 1
}

func square(x *float64) {
    *x = *x * *x
}

func swap(x, y *int) {
    aux := new(int)
    *aux = *x
    *x = *y
    *y = *aux
}
