package main

import (
    "strings"
    "net"
    "os"
    "os/exec"
    "bytes"
	"io"
    "github.com/donomii/svarmrgo"
    "fmt"
)

var history [3]string
var hpointer int
var lastNote string


func runCommand (cmd *exec.Cmd, stdin io.Reader) string{
    fmt.Println()
    fmt.Println("Started command")
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stderr = &err
	res := cmd.Run()
    fmt.Printf("Command result: %v\n", res)
	ret := fmt.Sprintf("%s\n%s", out, err)
    fmt.Println(ret)
    return ret
}

func l (conn net.Conn, k string) {
            cmd := exec.Command("c:\\Program Files\\AutoHotkey\\AutoHotkey","sendkey.ahk", k)
            runCommand(cmd,  strings.NewReader(""))
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "noteProcessor"})
            case "shutdown" :
            os.Exit(0)

         case "note" :
            switch m.Arg {
                case "A2" :
                    l(conn, "h")
                case "A#2" :
                    l(conn, "j")
                case "B2" :
                    l(conn, "k")
                case "C3" :
                    l(conn, "l")
                case "C#3" :
                    l(conn, "i")
            }
    }
}



func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
