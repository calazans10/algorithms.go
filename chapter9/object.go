package main

import "fmt"

type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
}

func main() {
    a := new(Android)
    a.Person.Name = "R2-D2"
    a.Model = "2"
    a.Talk()
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

