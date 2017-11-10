package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os/exec"
	//"strings"
	"github.com/donomii/svarmrgo"
)

type Message struct {
	Selector string
	Arg      string
}

func runCommand(cmd *exec.Cmd, stdin io.Reader) bytes.Buffer {
	cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out
}

func handleMessage(conn net.Conn, m svarmrgo.Message) {
	switch m.Selector {
	case "reveal-yourself":
		m.Respond(svarmrgo.Message{Selector: "announce", Arg: "message monitor"})
	default:
		log.Printf("%v:%v:%v\n", m.Selector, m.Arg, m.NamedArgs)
	}
}

func main() {
	conn := svarmrgo.CliConnect()
	svarmrgo.HandleInputs(conn, handleMessage)
}
