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
var serverConn net.Conn  //The connection to the server that started this module
var netPorts []net.Conn  //The remote server
func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
		out := []svarmrgo.Message{}
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
							m.Respond(svarmrgo.Message{Selector: "announce", Arg: relayID})
						case "shutdown" :
							shutdown = 1
						default:
							for _, v := range netPorts {
								now := time.Now()
								duration := time.Second * time.Duration(20)
								future := now.Add(duration)
								v.SetWriteDeadline(future)
								fmt.Fprintf(v, svarmrgo.WireFormat(m))
							}
						}
					return out
    }

func copyFromStdin(target net.Conn) {
	c1 := os.Stdin

    r := bufio.NewReader(c1)
    for {
        //svarmrgo.Debug("Outer copylines loop")
        l,err := r.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
	    fmt.Fprintf(target, fmt.Sprintf("%s", l))
    }
}
	
//FIXME this needs to handle multiple clients properly
func copyToStdout(source net.Conn) {
	
	c2 := os.Stdout

    r := bufio.NewReader(source)
    for {
        //svarmrgo.Debug("Outer copylines loop")
        l,err := r.ReadString('\n')
		if err != nil {
			return
		}
	    fmt.Fprintf(c2, fmt.Sprintf("%s", l))
    }
}

func announceMe (conn net.Conn) {
	for {
		//svarmrgo.Debug("Outer announce loop")
		svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "announce", Arg: relayID})
		time.Sleep(5000 * time.Millisecond)
	}
}




func main() {
	serverConn = svarmrgo.CliConnect()
	svarmrgo.HandleInputLoop(serverConn, handleMessage)
    // Listen for incoming connections.
    l, err := net.Listen("tcp", "0.0.0.0" + ":" + "4816")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    //fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	netPorts = append(netPorts, conn)
	go copyToStdout(conn)
}

