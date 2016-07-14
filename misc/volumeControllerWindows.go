package main
import (
    "net"
    "os/exec"
    "bytes"
    "fmt"
    "strings"
    "github.com/donomii/svarmrgo"
)


func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
	 case "reveal-yourself" :
	    response := svarmrgo.Message{Selector: "announce", Arg: "Windows volume control"}
	    m.Respond(response)
	 case "set-volume" :
		cmd := exec.Command("osascript",  "-e", fmt.Sprintf("set volume output volume %v --100%", m.Arg))
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()

	case "volume-up" :
		cmd := exec.Command("AutoHotkey", "volumeUp.ahk")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
	case "volume-down" :
		cmd := exec.Command("AutoHotkey", "volumeDown.ahk")
		//cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
	 case "mute" :
		cmd := exec.Command("osascript",  "-e", "set volume with output muted")
		//cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
	  case "unmute" :
		cmd := exec.Command("osascript",  "-e", "set volume without output muted")
		//cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()

    }
}

func main() {
	conn := svarmrgo.CliConnect()
        svarmrgo.HandleInputs(conn, handleMessage)
}
