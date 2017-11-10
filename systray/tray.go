package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"

	"github.com/donomii/svarmrgo"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
)

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	message := svarmrgo.Message{Selector: "announce", Arg: "systray"}
	out := []svarmrgo.Message{}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, message)

	case "tray-title":
		systray.SetTitle(m.Arg)
	case "tray-tooltip":
		systray.SetTooltip(m.Arg)
	case "tray-item":
		item := systray.AddMenuItem(m.Arg, m.Arg)
		go func() {
			for {
				_ = <-item.ClickedCh
				//mChange.SetTitle("I've Changed")
				svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "TrayItemClick", Arg: m.Arg})
			}
		}()
	}
	return out
}

func main() {
	runtime.GOMAXPROCS(4)
	conn := svarmrgo.CliConnect()
	go svarmrgo.HandleInputLoop(conn, handleMessage)

	onExit := func() {
		fmt.Println("Starting onExit")
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
		fmt.Println("Finished onExit")
	}
	// Should be called at the very beginning of main().
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetIcon(icon.Data)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Tooltray title")
		item := systray.AddMenuItem("Reveal Modules", "Force all modules to announce themselves")
		item1 := systray.AddMenuItem("Launch Modules", "Show module launcher")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItem("Unchecked", "Check Me")
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		systray.AddMenuItem("Ignored", "Ignored")
		mUrl := systray.AddMenuItem("Open Lantern.org", "my home")
		mQuit := systray.AddMenuItem("ÚÇÇÕç║", "Quit the whole app")
		systray.AddSeparator()
		mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")

		go func() {
			for {
				_ = <-item.ClickedCh
				//mChange.SetTitle("I've Changed")
				svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "reveal-yourself", Arg: ""})
			}
		}()

		go func() {
			for {
				_ = <-item1.ClickedCh
				//mChange.SetTitle("I've Changed")
				svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "start-module", Arg: "gui/gui"})
			}
		}()

		shown := true
		svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "ModuleStart", Arg: "systray"})
		svarmrgo.SendMessage(nil, svarmrgo.Message{Selector: "TrayStart", Arg: "systray"})

		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-mUrl.ClickedCh:
				open.Run("https://www.getlantern.org")
			case <-mToggle.ClickedCh:
				if shown {
					mQuitOrig.Hide()
					mEnabled.Hide()
					shown = false
				} else {
					mQuitOrig.Show()
					mEnabled.Show()
					shown = true
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
