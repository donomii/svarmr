package main

import (
	"github.com/andlabs/ui"
	"github.com/donomii/svarmrgo"
)

var box *ui.Box
var greeting *ui.Label

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	message := svarmrgo.Message{Selector: "announce", Arg: "Module Launcher"}
	out := []svarmrgo.Message{}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, message)
	case "gui-label":
		ui.QueueMain(func() {
			box.Append(ui.NewLabel(m.Arg), false)
		})
	case "tick":
	default:
		greeting.SetText("(" + m.Selector + ")" + m.Arg)
	}
	return out
}

func main() {
	conn := svarmrgo.CliConnect()
	go svarmrgo.HandleInputLoop(conn, handleMessage)
	svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "ModuleStart", Arg: "gui"})
	svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "GuiStart", Arg: "gui"})

	err := ui.Main(func() {
		name := ui.NewEntry()
		button := ui.NewButton("Launch")
		greeting = ui.NewLabel("")
		box = ui.NewVerticalBox()
		box.Append(ui.NewLabel("Module path:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Module Launcher", 200, 100, false)
		window.OnClosing(func(*ui.Window) bool {
			svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "ModuleQuit", Arg: "Module Launcher"})
			ui.Quit()
			return true
		})
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "start-module", Arg: name.Text()})
			//greeting.SetText("Hello, " + name.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
