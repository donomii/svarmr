package main

import (
  
    "github.com/donomii/svarmrgo"
    "time"

)

func handleMessage(m svarmrgo.Message) (out []svarmrgo.Message) {
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, svarmrgo.Message{Selector: "announce", Arg: "heartbeat"})

	}
	return
}

func main() {
    conn := svarmrgo.CliConnect()
	interval := 1
    m := svarmrgo.Message{ Selector: "heartbeat", Arg: "bompf"}
    go svarmrgo.HandleInputLoop(conn, handleMessage)
    for {
	    svarmrgo.SendMessage(conn, m)
	    time.Sleep(time.Duration(interval) * 1000 * time.Millisecond)
	}
}
