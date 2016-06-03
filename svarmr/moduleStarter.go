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
    "fmt"
)



func runCommand (cmd *exec.Cmd, stdin io.Reader) bytes.Buffer{
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stderr = &err
	cmd.Run()
    fmt.Printf(out.String())
    fmt.Printf(err.String())
    fmt.Printf("Command complete\n")
	return out
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
               svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "module starter"})
         case "shutdown" :
               os.Exit(0)
         case "start-module" :
               cmd := exec.Command(fmt.Sprintf("./%v", m.Arg), "localhost", "4816", "aaaaaa", "bbbbbb")
               fmt.Printf("%v", cmd)
               go runCommand(cmd,  strings.NewReader(""))
    }
}



func main() {
		conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }
