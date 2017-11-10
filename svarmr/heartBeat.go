package main

import (
	"net"
	"time"

	"github.com/donomii/svarmrgo"
)

func handleMessage(conn net.Conn, m svarmrgo.Message) {
	switch m.Selector {
	case "reveal-yourself":
		svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "announce", Arg: "heartbeat"})

	}
}

func main() {
	m := svarmrgo.Message{Selector: "heartbeat", Arg: "Hello World1"}
	for {
		svarmrgo.SendMessage(nil, m)
		time.Sleep(1000 * time.Millisecond)
	}
}
