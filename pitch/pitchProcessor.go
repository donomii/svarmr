package main

import (
    "strconv"
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

func closeEnough(a, b float64) bool {
    high := 1.05*a
    low  := 0.95*a
    if (low<b && b<high) {
        return true
    } else {
        return false
    }
}

func checkNote(conn net.Conn, pitch, noteHz float64, s string ) {
    if closeEnough(pitch, noteHz) {
        svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "pitch", Arg: s})
        history[hpointer] = s
        hpointer ++
        if hpointer>2 {
            hpointer = 0
        }
        if history[0]==history[1] && history[0]==history[2] {
            if lastNote != s {
                svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "note", Arg: s})
                lastNote = s
            }
        }
    }
}

func checkOctaves(conn net.Conn, pitch, noteHz float64, s string ) {
    for i:=8; i>0; i-- {
        //fmt.Printf("Comparing %v to %v\n", pitch,noteHz)
        checkNote(conn, pitch, noteHz, fmt.Sprintf("%v%v", s, i))
        noteHz=noteHz/2.0
    }
}


func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "mDnsProcessor"})
            case "shutdown" :
            os.Exit(0)

         case "pitch-detect" :
            pitch,_ := strconv.ParseFloat(m.Arg, 64)
            if pitch != 0.0 {
                //fmt.Printf(">Pitch: %v\n", pitch)
                checkOctaves(conn, pitch, 4186.01, "C")
                checkOctaves(conn, pitch, 4434.92, "C#")
                checkOctaves(conn, pitch, 4698.63, "D")
                checkOctaves(conn, pitch, 4978.03, "D#")
                checkOctaves(conn, pitch, 5274.04, "E")
                checkOctaves(conn, pitch, 5587.65, "F")
                checkOctaves(conn, pitch, 5919.91, "F#")
                checkOctaves(conn, pitch, 6271.93, "G")
                checkOctaves(conn, pitch, 6644.88, "G#")
                checkOctaves(conn, pitch, 7040.00, "A")
                checkOctaves(conn, pitch, 7458.62, "A#")
                checkOctaves(conn, pitch, 7902.13, "B")
                checkNote(conn,pitch, 82.41, "E2")
                checkNote(conn,pitch, 110.00, "A2")
                checkNote(conn,pitch, 146.83, "D3")
                checkNote(conn,pitch, 196.00, "G3")
                checkNote(conn,pitch, 246.94, "B3")
                checkNote(conn,pitch, 329.63, "E4")
                //fmt.Printf(">Finished checking\n")
            }
    }
}



func main() {
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
