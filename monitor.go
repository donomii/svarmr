package main
import (
    "net"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
	"io"
//"strings"
    "github.com/donomii/svarmrgo"
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

func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    switch m.Selector {
                         case "reveal-yourself" :
                            respondWith(conn, Message{Selector: "announce", Arg: "system monitor"})
                        default :
                            fmt.Printf("%v:%v\n", m.Selector, m.Arg)
                    }
                }

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
