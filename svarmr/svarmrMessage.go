//Send a svarmr message
//
//Svarmr messages are the internal system that makes svarmr work.  Normally, messages are created and sent by programs.  svarmrMessage allows you to create and sent these messages, allowing you to directly control svarmr.
//
//Examples:
//
//    svarmrMessage localhost 4816 user-notify "Hi there"
//
//will cause svarmr to pop up a hello message.
package main
import (
    "os"
    "github.com/donomii/svarmrgo"
)

func main() {
    conn := svarmrgo.CliConnect()
    selector := os.Args[3]
    arg := os.Args[4]
    m := svarmrgo.Message{ Selector: selector, Arg: arg}
    svarmrgo.SendMessage(conn, m)
}
