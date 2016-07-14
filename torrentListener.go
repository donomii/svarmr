package main
import (
    "github.com/donomii/svarmrgo"
    "net"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
    "time"
    "strings"
)

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
                response := svarmrgo.Message{Selector: "announce", Arg: "torrent control"}
                m.Respond(response)
         case "show-torrents" :
                cmd := exec.Command("deluge-console", "info")
                cmd.Stdin = strings.NewReader("")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()
                time.Sleep(5000 * time.Millisecond)
                fmt.Printf("%V\n", out)
                response := svarmrgo.Message{Selector: "torrent-status", Arg: string(out.Bytes())}
                m.Respond(response)
        case "add-torrent" :
                cmd := exec.Command("deluge-console", "add", m.Arg)
                cmd.Stdin = strings.NewReader("")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()
                response := svarmrgo.Message{Selector: "user-notify", Arg: "Started torrent"}
                m.Respond(response)
    }
}

func main() {
    conn, err := net.Dial("tcp", "localhost:4816")
    if err != nil {
        // handle error
    }
    for {
        if err != nil {
            fmt.Printf("Could not connect to server!")
        }
 svarmrgo.HandleInputs(conn, handleMessage)
    }
}
