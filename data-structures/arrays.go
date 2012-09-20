package main

import (
    "fmt"
)

// comprehension
func eachFloat32(x []float32, f func(int, float32) float32) []float32 {
    y := make([]float32, len(x))
    for i := 0; i < len(x); i++ {
        y[i] = f(i, x[i])
    }
    return y
}

func each(x []interface{}, f func(int, interface{}) interface{} ) []interface{} {

    y := make([]interface{}, len(x))

    for i := 0; i < len(x); i++ {
        y[i] = f(i, x[i])
    }

    return y
}

func main() {

    // warming up to arrays

    x := []int{1,2,3,4,5,6,7,8,9,0}
    fmt.Println("len(x) =", len(x))
    fmt.Println(x)
    
    y := x[3:6]
    fmt.Println("y =", y)
    
    y[0] = -1;
    fmt.Println("After changing y[0], x =", x)


    // heterogenous arrays
    type A struct {
        intValue int
    }

    type B struct {
        floatValue float32
    }
    type Variant interface{}

    a := A{12}
    b := B{1.2}
    fmt.Println(a,b)

    z := []Variant{a, b}
    fmt.Println(z)

    for i:=0; i < len(z); i++ {
        v := z[i]
        fmt.Println(i,v)
        switch t := v.(type) {
        case A: fmt.Println("A", t.intValue)
        case B: fmt.Println("B", t.floatValue)
        default:
            fmt.Println("not A nor B", t)
        }
    }

    z2 := eachFloat32([]float32{1.0,2.0}, func(i int, x float32) float32 {
        return 2*x / 3;
    })
    fmt.Println(z2)

    z3 := each([]interface{}{1,2,3,4}, func(i int, x interface{}) interface{} {return x.(int)+10})
    fmt.Println("z3", z3)
}