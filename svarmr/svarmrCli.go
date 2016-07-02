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
"github.com/tiborvass/uniline"
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


func cli(conn net.Conn) {
    prompt := "> "
    scanner := uniline.DefaultScanner()
    scanner.LoadHistory("~/.svarmrCli.history");
    for scanner.Scan(prompt) {
        line := scanner.Text()
        if len(line) > 0 {
            scanner.AddToHistory(line)
            cmd := strings.Split(line, " ")
            if cmd[0] == "q" || cmd[0] == "quit" || cmd[0] == "exit" {
                err := scanner.SaveHistory("~/.svarmrCli.history");
                fmt.Println(err)
                os.Exit(0)
            }
            if cmd[0] == "b" || cmd[0] == "broadcast" {
                selector := cmd[1]
                arg := cmd[2]
                msg := svarmrgo.Message{Selector: selector, Arg: arg}
                fmt.Printf("Sending: %v\n", msg)
                svarmrgo.SendMessage(conn, msg)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
               m.Respond(conn, svarmrgo.Message{Selector: "announce", Arg: "CommandLineInterface"})
         case "shutdown" :
               os.Exit(0)
         case "announce" :
            fmt.Println(m)
        default:
            //fmt.Println(m)
    }
}



func main() {
		conn := svarmrgo.CliConnect()
        go cli(conn)
        svarmrgo.HandleInputs(conn, handleMessage)
    }
