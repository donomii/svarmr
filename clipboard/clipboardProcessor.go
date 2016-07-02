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
	"regexp"
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
                               m.Respond(conn, svarmrgo.Message{Selector: "announce", Arg: "user notifier"})
                         case "shutdown" :
                               os.Exit(0)
                         case "clipboard-change" :
								if match, _ := regexp.MatchString("magnet", m.Arg); match {
									m.Respond(conn, svarmrgo.Message{Selector: "add-torrent", Arg: m.Arg})
								} else {
									m.Respond(conn, svarmrgo.Message{Selector: "user-notify-error", Arg: m.Arg})
								}
                    }
                }



func main() {
		conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }
