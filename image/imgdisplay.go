package main

import (
    "strings"
    "io/ioutil"
    "net"
    "os"
    "os/exec"
    "bytes"
    "github.com/donomii/svarmrgo"
    "fmt"
    "encoding/base64"
)

var history [3]string
var hpointer int
var lastNote string


func quickCommand (cmd *exec.Cmd) string{
    fmt.Println()
    fmt.Println("Started command")
    in := strings.NewReader("")
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stderr = &err
	res := cmd.Run()
    fmt.Printf("Command result: %v\n", res)
	ret := fmt.Sprintf("%s\n%s", out.String(), err.String())
    //fmt.Println(ret)
    return ret
}

func doDisplay (conn net.Conn, m svarmrgo.Message ) {
            pic, _ := base64.StdEncoding.DecodeString(m.Arg)
            ioutil.WriteFile("/tmp/temp_picture_for_display.jpg", pic, 0777)
            cmd := exec.Command("/usr/bin/open", "/tmp/temp_picture_for_display.jpg")
            quickCommand(cmd)
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "noteProcessor"})
            case "shutdown" :
            os.Exit(0)
         case "snapshot" :
            doDisplay(conn, m)
         case "image" :
            doDisplay(conn, m)
    }
}

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
