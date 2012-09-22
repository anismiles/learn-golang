package main
import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    echoer := func(name string, msg string, c chan string) {
        for i:=0; i < 10; i++ {
            fmt.Printf("\"%s\": %s\n", name, msg)
            n := time.Duration(50 + rand.Int() % 200)
            time.Sleep(n * time.Millisecond)
        }
        c <- name
    }

    c := make(chan string)
    for i:=0; i < 4; i++ {
        fmt.Println("\t\tStarting...", i)
        go func() {
            echoer(fmt.Sprintf("Echoer<%d>", i), "Hello", c);
        }()
    }

    for i:=0; i < 4; i++ {
        s:= <- c
        fmt.Printf("\t\t\"%s\" is done.\n", s)
    }
}
