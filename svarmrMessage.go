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
    out, _ := json.Marshal(m)
    fmt.Printf("%s\n", out)
    svarmrgo.RespondWith(conn, m)
}
