package main

import (
    "time"
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
            m.Respond(conn, svarmrgo.Message{Selector: "announce", Arg: "noteProcessor"})
            case "shutdown" :
            os.Exit(0)
    }
}

func watchOutput (conn net.Conn, b *bytes.Buffer) {
    for {
            line, err := b.ReadString('\n')
            if err != nil && err.Error() != "EOF" {
                fmt.Println(err)
            }
            if (len(line)>0) {
                svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "fs-change", Arg: line})
            } else {
                time.Sleep(1000 * time.Millisecond)
            }
    }
}

func doStartProc(conn net.Conn, cmd *exec.Cmd, in *strings.Reader, out, err bytes.Buffer) {
    cmd.Stdin = in
	cmd.Stdout = &out
	cmd.Stderr = &err
    go watchOutput(conn, &out)  //FIXME Move to startProc to not leak threads
	cmd.Run()
    fmt.Printf("Restarting: %v\n", cmd)
    doStartProc(conn, cmd, in, out, err)
}


func startProc(conn net.Conn, cmd *exec.Cmd) {
	var out bytes.Buffer
	var err bytes.Buffer
    in := strings.NewReader("")
    doStartProc(conn, cmd, in, out, err)
}



func main() {
    conn := svarmrgo.CliConnect()
    cmd := exec.Command("/usr/local/bin/fswatch", "/")
    go startProc(conn, cmd)
    svarmrgo.HandleInputs(conn, handleMessage)
}
