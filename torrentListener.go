package main
import (
    "net"
    "bufio"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
    "github.com/donomii/svarmrgo"
    "time"
    "strings"
)

func handleMessage (conn net.Conn, m svarmrgo.Message) {

                    switch m.Selector {
                         case "reveal-yourself" :
                            response := svarmrgo.Message{Selector: "announce", Arg: "torrent control"}
                            out, _ := json.Marshal(response)
                            fmt.Printf("%s\r\n", out)
                            fmt.Fprintf(conn, fmt.Sprintf("%s\r\n", out))
                         case "show-torrents" :
                                cmd := exec.Command("deluge-console", "info")
                                cmd.Stdin = strings.NewReader("")
                                var out bytes.Buffer
                                cmd.Stdout = &out
                                cmd.Run()
				time.Sleep(5000 * time.Millisecond)
				fmt.Printf("%V\n", out)
                                response := svarmrgo.Message{Selector: "torrent-status", Arg: string(out.Bytes())}
                                o, _ := json.Marshal(response)
                                fmt.Printf("%s\r\n", o)
                                fmt.Fprintf(conn, fmt.Sprintf("%s\r\n", o))
                        case "add-torrent" :
                                cmd := exec.Command("deluge-console", "add", m.Arg)
                                cmd.Stdin = strings.NewReader("")
                                var out bytes.Buffer
                                cmd.Stdout = &out
                                cmd.Run()
                                response := svarmrgo.Message{Selector: "user-notify", Arg: "Started torrent"}
				svarmrgo.RespondWith(conn, response)
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
