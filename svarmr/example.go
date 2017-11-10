package main

import (
	"fmt"
	"time"

	"github.com/donomii/svarmrgo"
)

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	out := []svarmrgo.Message{}
	//m3 := svarmrgo.Message{Selector: "systray-item", Arg: fmt.Sprintf("%v", "Clickme")}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, svarmrgo.Message{Selector: "announce", Arg: "example"})
	case "TrayStart":
		out = append(out, svarmrgo.Message{Selector: "tray-tooltip", Arg: "Svarmr controls"})
		out = append(out, svarmrgo.Message{Selector: "tray-title", Arg: "Svarmr"})
		out = append(out, svarmrgo.Message{Selector: "tray-item", Arg: "Launch Modules"})
		out = append(out, svarmrgo.Message{Selector: "tray-item", Arg: "Reveal Modules"})
		out = append(out, svarmrgo.Message{Selector: "debug-item", Arg: "Adding item"})
		//case "clock":
		//out = append(out, m3)
	case "TrayItemClick":
		if m.Arg == "Module Launcher" {
			out = append(out, svarmrgo.Message{Selector: "start-module", Arg: "gui/gui"})
		}

		if m.Arg == "Reveal Modules" {
			out = append(out, svarmrgo.Message{Selector: "reveal-yourself", Arg: ""})
		}

	}
	return out
}

func main() {
	conn := svarmrgo.CliConnect()
	svarmrgo.HandleInputLoop(conn, handleMessage)
	//for {

	m3 := svarmrgo.Message{Selector: "systray-item", Arg: fmt.Sprintf("%v", "Clickme")}
	time.Sleep(3 * time.Second)
	svarmrgo.SendMessage(nil, m3)
	time.Sleep(3 * time.Second)
	svarmrgo.SendMessage(nil, m3)
	time.Sleep(3 * time.Second)
	svarmrgo.SendMessage(nil, m3)

	//}
}
