//Pop up a message on the desktop
//
//svarmrMessage puts a temporary message on the screen, in Linux, MacOSX or Window.
//
//svarmrMessage sends the emssage through the native message system on each platform, so your messages will appear normal to the user.
package main

// This module requires the Notifu utility from http://www.paralint.com/projects/notifu/index.html#Download
// Copy it into a sub-directory under the server module, called "notifu"
//
// Messages:
//
//    user-notify: Displays the message stored in Arg
//    user-notify-error: Displays the message with an error title and icon, where possible
import (
	"bytes"
	"io"
	"log"
	"os/exec"

	"github.com/donomii/svarmrgo"
)

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	message := svarmrgo.Message{Selector: "announce", Arg: "echo"}
	switch m.Selector {
	case "reveal-yourself":
		m.Respond(message)
	}
	return []svarmrgo.Message{message}
}

func main() {
	conn := svarmrgo.CliConnect()
	svarmrgo.HandleInputLoop(conn, handleMessage)
}
