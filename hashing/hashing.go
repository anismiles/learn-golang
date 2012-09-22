package main
import (
    "fmt"
    "hash/adler32"
    "crypto/md5"
)

func main() {
    var s string = "hello world"
    fmt.Printf("s=\"%s\"\n", s)
    adler := adler32.New()
    adler.Write([]byte(s))
    fmt.Printf("adler32(s)=%x\n", adler.Sum32())
    md := md5.New()
    md.Write([]byte(s))
    v := uint64(md.Sum(nil))
    fmt.Printf("md5(s)=%x\n", v)
}
