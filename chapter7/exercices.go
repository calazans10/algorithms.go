package main

import "fmt"

func main() {
    fmt.Println(half(1))
    fmt.Println(half(2))

    list := []int{1,2,4,90,11}
    fmt.Println(max(list))
}

func half(number int) (int, bool) {
    if number % 2 == 0 {
        return number, true
    } else {
        return number, false
    }
}

func max(list []int) int {
    max := list[0]
    for _, value := range list {
        if max < value {
            max = value
        }
    }
    return max
}
