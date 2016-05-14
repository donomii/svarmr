package main

// This module requires the Notifu utility from http://www.paralint.com/projects/notifu/index.html#Download
// Copy it into a sub-directory called "notifu"
import (
    "net"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
	"io"
    "github.com/donomii/svarmrgo"
	"regexp"
)



func runCommand (cmd *exec.Cmd, stdin io.Reader) bytes.Buffer{
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out
}

func respondWith(conn net.Conn, response svarmrgo.Message) {
	out, _ := json.Marshal(response)
	fmt.Fprintf(conn, fmt.Sprintf("%s\r\n", out))
}


func handleMessage (conn net.Conn, m svarmrgo.Message) {
                    fmt.Printf("%v", m)
                    switch m.Selector {
                         case "reveal-yourself" :
								respondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "user notifier"})
                         case "clipboard-change" :
								if match, _ := regexp.MatchString("magnet", m.Arg); match {
									respondWith(conn, svarmrgo.Message{Selector: "add-torrent", Arg: m.Arg})
								} else {
									respondWith(conn, svarmrgo.Message{Selector: "user-notify-error", Arg: m.Arg})
								}
                    }
                }
            
    

func main() {
		conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }

