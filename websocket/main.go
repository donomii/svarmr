package main

import (
    "net"
    "fmt"
    "github.com/donomii/svarmrgo"
	"log"
	"net/http"

	"./chat"
)

var server *chat.Server

func main() {
    conn := svarmrgo.CliConnect()
    go svarmrgo.HandleInputs(conn, handleMessage)
	log.SetFlags(log.Lshortfile)

	// websocket server
	server = chat.NewServer("/entry")
    chat.Conn = conn
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    server.SendAll(&chat.Message{Author: m.Selector, Body: m.Arg})
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "system monitor"})
        default :
            fmt.Printf("%v:%v:%v\n", m.Selector, m.Arg, m.NamedArgs)
    }
}


