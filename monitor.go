package main
import (
    "net"
    "bufio"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
    "time"
	"io"
"strings"
)

type Message struct {
    Selector string
    Arg string
}

func runCommand (cmd *exec.Cmd, stdin io.Reader) bytes.Buffer{
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out
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
                var text = l
                fmt.Printf("%v\n", text)
                var m Message
                err := json.Unmarshal([]byte(text), &m)
                if err != nil {
                    fmt.Println("error:", err)
                } else {
                    fmt.Printf("%v", m)
                    switch m.Selector {
                         case "reveal-yourself" :
			        respondWith(conn, Message{Selector: "announce", Arg: "system monitor"})
                         case "show-processes" :
                                cmd := exec.Command("ps", "auxc")
				out := runCommand(cmd,  strings.NewReader(""))
                                respondWith(conn, Message{Selector: "process-list", Arg: string(out.Bytes())})
                    }
                }
            }
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
        handleConnection(conn)
    }
}
