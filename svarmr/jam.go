//Testing the server response when a client stops processing messages
package main
import (
    "net"
    "time"
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
                            m.Respond(svarmrgo.Message{Selector: "announce", Arg: "Testing jams"})
                        default :
                            //Emulate a client that has hung
                            for {
                                time.Sleep(1000 * time.Millisecond)
                            }
                    }
                }

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
