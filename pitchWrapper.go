package main

import (
    "net"
    _"fmt"
    "os/exec"
    "bytes"
	"io"
    "github.com/donomii/svarmrgo"
    "bufio"
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
                    }
                }



func main() {
    conn := svarmrgo.CliConnect()
    go svarmrgo.HandleInputs(conn, handleMessage)
    cmd := exec.Command("pitchDetect/pitchDetect")
    stdout, _ := cmd.StdoutPipe()
    cmd.Start()
    r := bufio.NewReader(stdout)
    for {
        line, _, _ := r.ReadLine()
        svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "pitch-detect", Arg: string(line)})
        //fmt.Println(string(line))
    }
}
