package main

import (
	"fmt"
    "bufio"
    "os"
)

func main() {
    f, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666);
    if err != nil {
        fmt.Println("Bad:", err)
        os.Exit(1)
    }
	fmt.Println("Good: %s", err)
    w := bufio.NewWriter(f)
    w.WriteString("Hello world\n")
    w.Flush()
    f.Close()
}
