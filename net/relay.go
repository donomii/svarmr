package main
import (

//"net/http"
//_"net/http/pprof"
//"runtime/pprof"
//"log"
    "net"
    "bufio"
    "fmt"
    "os/exec"
    "bytes"
    "time"
	"io"
    "github.com/donomii/svarmrgo"
	"os"
"flag"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

var shutdown int = 0

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
        svarmrgo.Debug("Outer handlemessage loop")
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
                        shutdown = 1
                    }
    }

func copyLines(c1, c2 net.Conn) {
    r := bufio.NewReader(c1)
    for {
        svarmrgo.Debug("Outer copylines loop")
        l,err := r.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
	    fmt.Fprintf(c2, fmt.Sprintf("%s", l))
    }
}

func announceMe (conn net.Conn) {
	for {
        svarmrgo.Debug("Outer announce loop")
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
	go announceMe(conn3)
	go svarmrgo.HandleInputs (conn3, handleMessage)
	fmt.Printf("Started relay\n")
//go func() {
    //log.Println(http.ListenAndServe("localhost:6060", nil))
//}()
    for {
	time.Sleep(5000 * time.Millisecond)

        if shutdown==1 {
        return
    }
    }
}

