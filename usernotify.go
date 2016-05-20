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
	"encoding/json"
)



func runCommand (cmd *exec.Cmd, stdin io.Reader) bytes.Buffer{
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out
}


type NotifyArgs struct {
    Message string
    Title string
    Level string
    Duration string
}


func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    switch m.Selector {
                         case "reveal-yourself" :
			        svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "user notifier"})
                         case "user-notify" :
                                cmd := exec.Command("notifu/notifu.exe", "/m", m.Arg, "/p", "Svarmr", "/t", "info")
								runCommand(cmd,  strings.NewReader(""))
                                cmd = exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%v\" with title \"Svarmr\" ", m.Arg))
								runCommand(cmd,  strings.NewReader(""))
						case "user-notify-error" :
                                cmd := exec.Command("notifu/notifu.exe", "/m", m.Arg, "/p", "Svarmr", "/t", "error")
								runCommand(cmd,  strings.NewReader(""))
                                cmd = exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%v\" with title \"Svarmr\" ", m.Arg))
								runCommand(cmd,  strings.NewReader(""))
						case "user-notify-custom" :
							    var a NotifyArgs
								err := json.Unmarshal([]byte(m.Arg), &a)
								if err != nil {
                                    fmt.Printf("Error: %v\n", err)
                                }
                                cmd := exec.Command("notifu/notifu.exe", "/m", a.Message, "/p", a.Title, "/t", a.Level, "/d", a.Duration)
								//fmt.Printf("%v\n",cmd)
								runCommand(cmd,  strings.NewReader(""))
                                cmd = exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%v\" with title \"Svarmr\" ", m.Arg))
								runCommand(cmd,  strings.NewReader(""))
                                //svarmrgo.RespondWith(conn, Message{Selector: "process-list", Arg: string(out.Bytes())})
                    }
                }



func main() {
		conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }
