package main

// This module requires the Notifu utility from http://www.paralint.com/projects/notifu/index.html#Download
// Copy it into a sub-directory called "notifu"
import (
    "net"
    "fmt"
    "os/exec"
    "bytes"
	"io"
	"strings"
    "github.com/donomii/svarmrgo"
)



func runCommand (cmd *exec.Cmd, stdin io.Reader) bytes.Buffer{
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out
}




func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    fmt.Printf("%v", m)
                    switch m.Selector {
                         case "reveal-yourself" :
			        svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "user notifier"})
                         case "user-notify" :
                                cmd := exec.Command("notifu/notifu.exe", "/m", m.Arg, "/p", "Svarmr", "/t", "info")
								 runCommand(cmd,  strings.NewReader(""))
						case "user-notify-error" :
                                cmd := exec.Command("notifu/notifu.exe", "/m", m.Arg, "/p", "Svarmr", "/t", "error")
								 runCommand(cmd,  strings.NewReader(""))
                                //svarmrgo.RespondWith(conn, Message{Selector: "process-list", Arg: string(out.Bytes())})
                    }
                }
            
    

func main() {
		conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }

