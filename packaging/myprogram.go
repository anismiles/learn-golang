package main

import (
    "fmt"
    X "example/newmath"
)

func main() {
    c := X.Sqrt(100.0)
    fmt.Printf("Hello %f\n", c)
    a, b := X.Updown(10)
    fmt.Printf("World %d %d\n", a, b)
}
