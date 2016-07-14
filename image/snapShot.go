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
	ret := fmt.Sprintf("%s\n%s", out, err)
    fmt.Println(ret)
    return ret
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            m.Respond(svarmrgo.Message{Selector: "announce", Arg: "noteProcessor"})
            case "shutdown" :
            os.Exit(0)
         case "take-picture" :
            cmd := exec.Command("/usr/local/bin/imagesnap", "-w", "1.00", "temp_picture.jpg")
            quickCommand(cmd)
            pic, _ := ioutil.ReadFile("temp_picture.jpg")
            enc_pic := base64.StdEncoding.EncodeToString(pic)
            m.Respond(svarmrgo.Message{Selector: "snapshot", Arg: enc_pic})
    }
}

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
