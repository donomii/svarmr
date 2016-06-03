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

func doRecognise (conn net.Conn, m svarmrgo.Message ) {
            pic, _ := base64.StdEncoding.DecodeString(m.Arg)
            ioutil.WriteFile("/tmp/temp_picture_for_recogniser.jpg", pic, 0777)
            cmd := exec.Command("aux/recognise.sh", "/tmp/temp_picture_for_recogniser.jpg")
            answer := quickCommand(cmd)
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "image-recognised", Arg: answer})
    
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "noteProcessor"})
            case "shutdown" :
            os.Exit(0)
         case "snapshot" :
            doRecognise(conn, m)
         case "image" :
            doRecognise(conn, m)
         case "recognise" :
            doRecognise(conn, m)
    }
}

func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
