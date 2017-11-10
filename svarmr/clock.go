package main

import (
	"fmt"
	"time"

	"github.com/donomii/svarmrgo"
)

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	out := []svarmrgo.Message{}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, svarmrgo.Message{Selector: "announce", Arg: "clock"})
	}
	return out
}

func main() {
	conn := svarmrgo.CliConnect()
	svarmrgo.HandleInputLoop(conn, handleMessage)
	for {
		now := time.Now()
		m := svarmrgo.Message{Selector: "tick", Arg: fmt.Sprintf("%v", now.Unix())}
		//m2 := svarmrgo.Message{Selector: "gui-label", Arg: fmt.Sprintf("%v", now.Unix())}
		//m3 := svarmrgo.Message{Selector: "systray-item", Arg: fmt.Sprintf("%v", now.Unix())}
		//svarmrgo.SendMessage(nil, m3)
		svarmrgo.SendMessage(nil, m)
		//svarmrgo.SendMessage(nil, m2)

		time.Sleep(1000 * time.Millisecond)
	}
}
