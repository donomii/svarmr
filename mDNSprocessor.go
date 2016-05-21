package main

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

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "mDnsProcessor"})
            case "shutdown" :
            os.Exit(0)

         case "mdns-found-ipv4" :
            /'nix
            arr := strings.Split(m.Arg, ":")
            cmd := exec.Command("./relay", os.Args[1], os.Args[2],arr[0], arr[1])
            fmt.Printf("Starting new relay %v\n", cmd)
            go runCommand(cmd, strings.NewReader("some input") )

            //Windows
            cmd = exec.Command("relay.exe", os.Args[1], os.Args[2],arr[0], arr[1])
            fmt.Printf("Starting new relay %v\n", cmd)
            go runCommand(cmd, strings.NewReader("some input") )
    }
}



func main() {
	conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
    }
