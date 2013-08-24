package main

import "fmt"

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))

    x, y := f()
    fmt.Println(x + y)

    fmt.Println(add(1,2,3))

    xt := []int{1,2,4}
    fmt.Println(add(xt...))
}

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func f() (int, int) {
    return 5, 6
}

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}
