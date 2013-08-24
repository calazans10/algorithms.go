package main

import "fmt"

// this is a comment

func main() {
    var x string
    x = "Hello World"
    fmt.Println(len(x))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
}
