package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {
    p1 := person{"Alice", 30}
    fmt.Println(p1)

    p2 := person{name: "Bob", age: 25}
    fmt.Println(p2)

    var p3 person
    p3.name = "Charlie"
    p3.age = 20
    fmt.Println(p3)

    p4 := &person{name: "David", age: 35}
    fmt.Println(p4)
}
