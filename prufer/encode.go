package main

import (
    "./tree"
    "fmt"
)

func main() {
    n := tree.Make(0)
    n.Label = 1
    fmt.Println(n)
}
