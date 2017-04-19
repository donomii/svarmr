package main
import (
    "strconv"
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
    if len(os.Args) < 4 {
        panic("Please supply a broadcast interval as the third command line argument")
    }
    arg := os.Args[3]
    interval, _ := strconv.Atoi(arg) 
    m := svarmrgo.Message{ Selector: "heartbeat", Arg: arg}
    go svarmrgo.HandleInputs(conn, handleMessage)
    for {
	    svarmrgo.SendMessage(conn, m)
	    time.Sleep(time.Duration(interval) * 1000 * time.Millisecond)
	}
}
