package main
import (
    "net"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
    "strings"
    "github.com/donomii/svarmrgo"
)

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
                response := svarmrgo.Message{Selector: "announce", Arg: "MacOSX volume control"}
                out, _ := json.Marshal(response)
                fmt.Printf("%sr\n", out)
                svarmrgo.RespondWith(conn, response)
         case "set-volume" :
                fmt.Printf("Volume up\n")
                cmd := exec.Command("osascript",  "-e", fmt.Sprintf("set volume output volume %v --100%", m.Arg))
                cmd.Stdin = strings.NewReader("some input")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()

        case "volume-up" :
                fmt.Printf("Volume up\n")
                cmd := exec.Command("osascript",  "-e", fmt.Sprintf("set volume output volume (output volume of (get volume settings) + %v) --100%", m.Arg))
                cmd.Stdin = strings.NewReader("some input")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()
        case "volume-down" :
                fmt.Printf("Volume down\n")
                cmd := exec.Command("osascript",  "-e", fmt.Sprintf("set volume output volume (output volume of (get volume settings) - %v) --100%", m.Arg))
                //cmd.Stdin = strings.NewReader("some input")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()
         case "mute" :
                fmt.Printf("Volume mute\n")
                cmd := exec.Command("osascript",  "-e", "set volume with output muted")
                //cmd.Stdin = strings.NewReader("some input")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()
          case "unmute" :
                fmt.Printf("Volume unmute\n")
                cmd := exec.Command("osascript",  "-e", "set volume without output muted")
                //cmd.Stdin = strings.NewReader("some input")
                var out bytes.Buffer
                cmd.Stdout = &out
                cmd.Run()

    }
}

func main() {
    conn, err := net.Dial("tcp", "localhost:4816")
    if err != nil {
        // handle error
    }
    for {
        if err != nil {
            fmt.Printf("Could not connect to server!")
        }
        svarmrgo.HandleInputs(conn, handleMessage)
    }
}
