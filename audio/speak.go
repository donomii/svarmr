package main

import (
    "net"
    "os"
    "os/exec"
    "github.com/donomii/svarmrgo"
)

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "announce", Arg: "speak"})
         case "shutdown" :
            os.Exit(0)
         case "say" :
            cmd := exec.Command("say", "-v", "Daniel",  m.Arg)
            go svarmrgo.QuickCommandStdout(cmd)
    }
}

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
