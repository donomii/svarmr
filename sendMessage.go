package main
import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/donomii/svarmrgo"
)

func main() {
    conn := svarmrgo.CliConnect()
    selector := os.Args[3]
    arg := os.Args[4]
    m := svarmrgo.Message{ Selector: selector, Arg: arg}
    fmt.Printf("%v\r\n", m)
    out, _ := json.Marshal(m)
    fmt.Printf("%s\r\n", out)
    fmt.Fprintf(conn, fmt.Sprintf("%s\r\n", out))
}
