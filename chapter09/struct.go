package main

import ("fmt"; "math")

type Shape interface {
    area() float64
}

type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func main() {
    c := Circle{0, 0, 5}

    area := circleArea(&c)
    fmt.Println(area)

    fmt.Println(c.area())

    r := Rectangle{0,0,10,10}
    s := Shape(&r)
    fmt.Println(s.area())
}

func circleArea(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
