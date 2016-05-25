package main
import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/donomii/svarmrgo"
    "time"
"net"
)


func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    switch m.Selector {
                         case "reveal-yourself" :
			                        svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "heartbeat"})
                    }
                }

func main() {
    conn := svarmrgo.CliConnect()
    arg := os.Args[3]
    m := svarmrgo.Message{ Selector: "heartbeat", Arg: arg}
    out, _ := json.Marshal(m)
    fmt.Printf("%s\n", out)
    go svarmrgo.HandleInputs(conn, handleMessage)
    for {
	    svarmrgo.RespondWith(conn, m)
	    time.Sleep(1000 * time.Millisecond)
	}
}
