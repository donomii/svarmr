package main
import (
    "net"
    "bufio"
    "fmt"
    "encoding/json"
    "os/exec"
    "bytes"
    "time"
    "strings"
)

type Message struct {
    Selector string
    Arg string
}

func handleConnection (conn net.Conn) {
    fmt.Sprintf("%V", conn)
    time.Sleep(500 * time.Millisecond)
    r := bufio.NewReader(conn)
    for {
        l,_ := r.ReadString('\n')
        if (l!="") {
                var text = l
                fmt.Printf("%v\n", text)
                var m Message
                err := json.Unmarshal([]byte(text), &m)
                if err != nil {
                    fmt.Println("error:", err)
                } else {
                    fmt.Printf("%v", m)
                    switch m.Selector {
                         case "reveal-yourself" :
                            response := Message{Selector: "announce", Arg: "Windows volume control"}
                            out, _ := json.Marshal(response)
                            fmt.Printf("%s\n", out)
                            svarmrgo.RespondWith(conn, out)
                         case "set-volume" :
                                fmt.Printf("Volume up\n")
                                cmd := exec.Command("osascript",  "-e", fmt.Sprintf("set volume output volume %v --100%", m.Arg))
                                cmd.Stdin = strings.NewReader("some input")
                                var out bytes.Buffer
                                cmd.Stdout = &out
                                cmd.Run()

                        case "volume-up" :
                                fmt.Printf("Volume up\n")
                                cmd := exec.Command("AutoHotkey", "volumeUp.ahk")
                                cmd.Stdin = strings.NewReader("some input")
                                var out bytes.Buffer
                                cmd.Stdout = &out
                                cmd.Run()
                        case "volume-down" :
                                fmt.Printf("Volume down\n")
                                cmd := exec.Command("AutoHotkey", "volumeDown.ahk")
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
            }
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
        handleConnection(conn)
    }
}
