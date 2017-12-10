package main
import (

//"net/http"
//_"net/http/pprof"
//"runtime/pprof"
//"log"
	"math/rand"
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
var netPort net.Conn  //The remote server

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
						case "svarmrConnect" :
							target := m.Arg
							relayID = fmt.Sprintf("relay (%v)",rand.Int(), os.Getpid())
							altRelayID = fmt.Sprintf("relay %v:%v - %v",rand.Int(), os.Getpid(), target)
							
							conn2 := svarmrgo.ConnectHub(target)
							go copyToStdout(conn2)
						default:
							if netPort != nil {
								duration := time.Second * time.Duration(20)
								now := time.Now()
								future := now.Add(duration)
								netPort.SetWriteDeadline(future)
								fmt.Fprintf(netPort, svarmrgo.WireFormat(m))
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
	
func copyToStdout(source net.Conn) {
	
	c2 := os.Stdout

    r := bufio.NewReader(source)
    for {
        //svarmrgo.Debug("Outer copylines loop")
        l,err := r.ReadString('\n')
		if err != nil {
			os.Exit(1)
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
	//go svarmrgo.HandleInputs (serverConn, handleMessage)
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

