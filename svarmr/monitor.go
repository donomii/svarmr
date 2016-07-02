package main
import (
    "net"
    "fmt"
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

func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    switch m.Selector {
                         case "reveal-yourself" :
                            m.Respond(conn, svarmrgo.Message{Selector: "announce", Arg: "system monitor"})
                        default :
                            fmt.Printf("%v:%v:%v\n", m.Selector, m.Arg, m.NamedArgs)
                    }
                }

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
