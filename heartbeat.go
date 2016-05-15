package main
import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/donomii/svarmrgo"
    "time"
)

func main() {
    conn := svarmrgo.CliConnect()
    arg := os.Args[3]
    m := svarmrgo.Message{ Selector: "heartbeat", Arg: arg}
    out, _ := json.Marshal(m)
    fmt.Printf("%s\r\n", out)
    for {
	    fmt.Fprintf(conn, fmt.Sprintf("%s\n", out))
	    time.Sleep(1000 * time.Millisecond)
	}
}
