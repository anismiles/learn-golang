package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("specify the file")
        os.Exit(1)
    }
    f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
    if err != nil {
        fmt.Printf("file \"%s\" cannot be opened.\n", os.Args[1])
        os.Exit(2)
    }
    r := bufio.NewReader(f)
    for err = nil; err != io.EOF; {
        var s []byte
        s, _, err = r.ReadLine() // err := causes a new 'err' to be made
        if err == nil {
            fmt.Printf("> \"%s\"\n", string(s))
        }
    }
}
