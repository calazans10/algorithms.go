package main

import (
    "fmt"
    "github.com/calazans10/math"
)

func main() {
    xs := []float64{1,2,3,4}
    fmt.Println("Avg:", math.Average(xs))
    fmt.Println("Sum:", math.Sum(xs))
    fmt.Println("Min:", math.Min(xs))
    fmt.Println("Max:", math.Max(xs))
}
