package main

import (
	"os"

	"github.com/andlabs/ui"
	"github.com/donomii/svarmrgo"
)

var box *ui.Box
var greeting *ui.Label

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	message := svarmrgo.Message{Selector: "announce", Arg: "gui"}
	out := []svarmrgo.Message{}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, message)
	case "tick":
	default:
		//greeting.SetText("(" + m.Selector + ")" + m.Arg)
	}
	return out
}

func main() {
	conn := svarmrgo.CliConnect()
	go svarmrgo.HandleInputLoop(conn, handleMessage)
	svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "ModuleStart", Arg: "userMessage"})
	svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "userMessage", Arg: "start"})

	err := ui.Main(func() {
		name := ui.NewEntry()
		arg := ui.NewEntry()
		button := ui.NewButton("Send")
		greeting = ui.NewLabel("")
		box = ui.NewVerticalBox()
		box.Append(ui.NewLabel("Selector"), false)
		box.Append(name, false)
		box.Append(ui.NewLabel("Arg"), false)
		box.Append(arg, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Send Message", 200, 100, false)
		window.OnClosing(func(*ui.Window) bool {
			svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "ModuleQuit", Arg: "userMessage"})
			ui.Quit()
			os.Exit(0)
			return true
		})
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: name.Text(), Arg: arg.Text()})
			//greeting.SetText("Hello, " + name.Text() + "!")
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
