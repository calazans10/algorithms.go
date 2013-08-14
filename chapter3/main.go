package main

import "fmt"

func main() {
    var v1 int64 = 321325
    var v2 int64 = 424521

    fmt.Println("1 + 1 =", 1.0 + 1.0)
    fmt.Println("321325 x 424521 =", v1 * v2)
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
    fmt.Println((true && false) || (false && true) || !(false && false))
}
