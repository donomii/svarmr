package main
import (
    "net"
    "bufio"
    "fmt"
    "time"
    "os"
    "encoding/json"
)

type Message struct {
    Selector string
    Arg string
}

func respondWith(conn net.Conn, response Message) {
out, _ := json.Marshal(response)
fmt.Fprintf(conn, fmt.Sprintf("%s\r\n", out))
}

func handleConnection (conn net.Conn) {
    fmt.Sprintf("%V", conn)
    time.Sleep(500 * time.Millisecond)
    r := bufio.NewReader(conn)
    for {
        l,_ := r.ReadString('\n')
        if (l!="") {
                var m Message
                json.Unmarshal([]byte(l), &m)
                switch m.Selector {
                    case "reveal-yourself":
                        respondWith(conn, Message{Selector: "announce", Arg: "debug monitor"})
                    default:
}
                fmt.Printf("%v\n", l)
		fmt.Printf("%v\n", m.Arg)

            }
        }
    }

func main() {
    server := os.Args[1]
    port := os.Args[2]
    conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", server, port))
    if err != nil {
        // handle error
    }
    for {
        if err != nil {
            fmt.Printf("Could not connect to server!")
        }
        handleConnection(conn)
    }
}
