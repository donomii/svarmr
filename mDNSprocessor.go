package main

// This module requires the Notifu utility from http://www.paralint.com/projects/notifu/index.html#Download
// Copy it into a sub-directory called "notifu"
import (
    "net"
"os"
    "os/exec"
    "bytes"
	"io"
    "github.com/donomii/svarmrgo"
    "strings"
)



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
								                   svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "mDnsProcessor"})
                                 case "shutdown" :
                                           os.Exit(0)

                         case "mdns-found-ipv4" :
				arr := strings.Split(m.Arg, ":")
				      cmd := exec.Command("relay", os.Args[1], os.Args[2],arr[0], arr[1])
				go runCommand(cmd, strings.NewReader("some input") )
				}
                    }



func main() {
	conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }
