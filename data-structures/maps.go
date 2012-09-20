package main
import (
    . "fmt"
    "io/ioutil"
    "os"
    "strings"
    "sort"
)

type MyMap map[string] int
type Entry struct {
    Key string
    Val int
}
type Entries []Entry
func makeEntries(myMap MyMap) Entries {
    entries := make(Entries, len(myMap))
    for k,v := range myMap {
        entries = append(entries, Entry{k,v})
    }
    return entries
}
func (list Entries) Swap(i, j int) {
    list[i], list[j] = list[j], list[i]
}
func (list Entries) Len() int {
    return len(list)
}
func (list Entries) Less(i, j int) bool {
    return list[i].Val < list[j].Val
}

func main() {
    filename := "maps.go"
    Println("opening file", filename)
    buf, err := ioutil.ReadFile(filename)

    if err != nil {
        Println("Not able to open file")
        os.Exit(1)
    }

    Println("buf len", len(buf))
    var fields []string
    fields = strings.Fields(string(buf))
    Printf("there are %d\n", len(fields))

    counts := make(MyMap)
    for _,field := range fields {
        i, ok := counts[field]
        if ! ok {
            counts[field] = 1
        } else {
            counts[field] = i+1
        }
    }

    // build the list of keys
    entries := makeEntries(counts)
    sort.Sort(entries)
    for i,e := range entries {
        Printf("%d: \"%s\" (%d)\n", i, e.Key, e.Val)
    }
}
