package main
import (
    "net"
    "bufio"
    "fmt"
    "os/exec"
    "bytes"
    "time"
	"io"
    "github.com/donomii/svarmrgo"
	"os"
)

func runCommand (cmd *exec.Cmd, stdin io.Reader) bytes.Buffer{
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out
}
var relayID string
var altRelayID string
func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    switch m.Selector {
                         case "announce" :
				if relayID == m.Arg { //We either have a loop or we have two copies of the same relay running
                    fmt.Printf("Detected routing loop - exiting!\n");
					os.Exit(0)
				}
				if altRelayID == m.Arg { //We either have a loop or we have two copies of the same relay running
                    fmt.Printf("Detected routing loop - exiting!\n");
					os.Exit(0)
				}
                         case "reveal-yourself" :
				svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: relayID})
              case "shutdown" :
                        os.Exit(0)
                    }
    }

func copyLines(c1, c2 net.Conn) {
    r := bufio.NewReader(c1)
    for {
        l,err := r.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
	    fmt.Fprintf(c2, fmt.Sprintf("%s\n", l))
    }
}

func announceMe (conn net.Conn) {
	for {
	svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: relayID})
	time.Sleep(5000 * time.Millisecond)
}
}

func main() {
    server1 := os.Args[1]
    port1 := os.Args[2]
    server2 := os.Args[3]
    port2 := os.Args[4]
    relayID = fmt.Sprintf("relay %v:%v - %v:%v (%v)",server1, port1, server2,port2, os.Getpid())
    altRelayID = fmt.Sprintf("relay %v:%v - %v:%v (%v)",server2, port2, server1,port1, os.Getpid())
    conn1 := svarmrgo.ConnectHub(server1, port1)
    conn2 := svarmrgo.ConnectHub(server2, port2)
	conn3 := svarmrgo.ConnectHub(server1, port1)
    go copyLines(conn1, conn2)
    go copyLines(conn2, conn1)
	go svarmrgo.HandleInputs (conn3, handleMessage)
	go announceMe(conn3)
	fmt.Printf("Started relay\n")
	for {}
}

