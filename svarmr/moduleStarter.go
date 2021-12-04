package main
//Starts another module, with correct command line opts

import (
    "net"
"os"
    "os/exec"
    "github.com/donomii/svarmrgo"
    "fmt"
)



func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
               svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "announce", Arg: "module starter"})
         case "shutdown" :
               os.Exit(0)
         case "start-module" :
               cmd := exec.Command(fmt.Sprintf("./%v", m.Arg), "localhost", "4816", "aaaaaa", "bbbbbb")
               fmt.Printf("%v", cmd)
               go svarmrgo.QuickCommandStderr(cmd)
    }
}



func main() {
		conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }
