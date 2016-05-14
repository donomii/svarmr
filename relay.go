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

func handleConnection (conn net.Conn) {
    fmt.Sprintf("%V", conn)
    time.Sleep(500 * time.Millisecond)
    r := bufio.NewReader(conn)
    for {
        l,_ := r.ReadString('\n')
        if (l!="") {
                var text = l
                fmt.Printf("%v\n", text)
                var m svarmrgo.Message
                err := json.Unmarshal([]byte(text), &m)
                if err != nil {
                    fmt.Println("error:", err)
                } else {
                    fmt.Printf("%v", m)
                    switch m.Selector {
                         case "reveal-yourself" :
								svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "system monitor"})
                         case "show-processes" :
                                cmd := exec.Command("ps", "auxc")
								out := runCommand(cmd,  strings.NewReader(""))
                                svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "process-list", Arg: string(out.Bytes())})
                    }
                }
            }
        }
    }

func copyLines(c1, c2 net.Conn) {
    r := bufio.NewReader(c1)
    for {
        l,err := r.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s", l)
	    fmt.Fprintf(c2, fmt.Sprintf("%s\n", l))
    }
}

func main() {
    server1 := os.Args[1]
    port1 := os.Args[2]
    server2 := os.Args[3]
    port2 := os.Args[4]
    conn1 := svarmrgo.ConnectHub(server1, port1)
    conn2 := svarmrgo.ConnectHub(server2, port2)
    go copyLines(conn1, conn2)
    go copyLines(conn2, conn1)
	fmt.Printf("Started relay\n")
	for {}
}