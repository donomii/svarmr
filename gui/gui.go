package main

import (
	"github.com/andlabs/ui"
	"github.com/donomii/svarmrgo"
)

var box *ui.Box

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	message := svarmrgo.Message{Selector: "announce", Arg: "gui"}
	out := []svarmrgo.Message{}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, message)
	case "gui-label":
		ui.QueueMain(func() {

			box.Append(ui.NewLabel(m.Arg), false)
		})
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
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box = ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.OnClosing(func(*ui.Window) bool {
			svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "ModuleQuit", Arg: "gui"})
			ui.Quit()
			return true
		})
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "InputEvent", Arg: name.Text()})
			greeting.SetText("Hello, " + name.Text() + "!")
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
