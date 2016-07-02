package main
import (
    "os"
    "github.com/donomii/svarmrgo"
    "time"
"net"
)


func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    switch m.Selector {
                         case "reveal-yourself" :
                            svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "announce", Arg: "heartbeat"})
                    }
                }

func main() {
    conn := svarmrgo.CliConnect()
    arg := os.Args[3]
    m := svarmrgo.Message{ Selector: "heartbeat", Arg: arg}
    go svarmrgo.HandleInputs(conn, handleMessage)
    for {
	    svarmrgo.SendMessage(conn, m)
	    time.Sleep(1000 * time.Millisecond)
	}
}
