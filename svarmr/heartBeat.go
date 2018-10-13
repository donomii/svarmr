package main

import (
    "strconv"
    "os"
    "github.com/donomii/svarmrgo"
    "time"
"log"
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
	arg := "1"
    if len(os.Args) < 4 {
        log.Println("Please supply a broadcast interval as the third command line argument")
    } else {
		arg = os.Args[3]
	}
    interval, _ := strconv.Atoi(arg) 
    m := svarmrgo.Message{ Selector: "heartbeat", Arg: arg}
    go svarmrgo.HandleInputLoop(conn, handleMessage)
    for {
	    svarmrgo.SendMessage(conn, m)
	    time.Sleep(time.Duration(interval) * 1000 * time.Millisecond)
	}
}
