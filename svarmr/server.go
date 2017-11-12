//The message hub.  Server relays messages between programs
//
//Start with
//
//    server localhost 4816
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/donomii/svarmrgo"
)

const queueLength int = 1

var inMessages int = 0
var outMessages int = 0

type connection struct {
	port   net.Conn
	handle *subProx
	raw    string
}

type subProx struct {
	In  io.WriteCloser
	Out io.ReadCloser
	Err io.ReadCloser
	Cmd *exec.Cmd
}

//Read incoming messages and place them on the input queue
func handleSubprocConnection(conn *subProx, Q chan connection) {

	reader := bufio.NewReader(conn.Out)

	for {
		t, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println("Client disconnected: ", err)
			return
		}
		var m connection = connection{nil, conn, t}
		Q <- m
	}

}

func StartSubproc(orig_cmd string, args []string) *subProx {
	//It turns out that cmd.Start() doesn't actually tell us if the subprocess started,
	//just that the internal call succeeded.  So we have no actual way of telling if the
	//subprocess started, without waiting for it to quit.
	var cmd string
	var handle *subProx
	log.Println("OS:", runtime.GOOS)
	if runtime.GOOS == "windows" {
		//This is fucking retarded
		detectPath := fmt.Sprintf("%s.bat", orig_cmd)
		cmd = strings.Replace(fmt.Sprintf("%s.bat", orig_cmd), "/", "\\", -1)
		log.Println("Trying ", cmd)
		if _, err := os.Stat(detectPath); !os.IsNotExist(err) {
			handle = ActualStartSubproc(cmd, args)
		} else {
			cmd = fmt.Sprintf("%s.exe", orig_cmd)
			log.Println("Trying ", cmd)
			handle = ActualStartSubproc(cmd, args)
			return handle
		}
	} else {
		cmd = fmt.Sprintf("%s.sh", orig_cmd)
		log.Println("Trying ", cmd)
		if _, err := os.Stat(cmd); !os.IsNotExist(err) {
			handle = ActualStartSubproc(cmd, args)
			return handle
		} else {
			log.Println("Trying ", orig_cmd)
			if _, err := os.Stat(orig_cmd); !os.IsNotExist(err) {
				handle := ActualStartSubproc(orig_cmd, args)
				if handle != nil {
					log.Println("Succeeded starting ", orig_cmd)
					return handle
				}
			}
		}
	}

	return nil
}

func ActualStartSubproc(cmd string, args []string) *subProx {
	grepCmd := exec.Command(cmd, args...)

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	//Err, _ := grepCmd.StderrPipe()

	err := grepCmd.Start()
	log.Println("Start command result:", err)
	if err != nil {
		return nil
	}
	p := subProx{grepIn, grepOut, nil, grepCmd}
	subprocList = append(subprocList, &p)
	go handleSubprocConnection(&p, inQ)
	inQ <- connection{nil, nil, svarmrgo.WireFormat(svarmrgo.Message{Selector: "user-notify", Arg: "Service started: " + cmd})}
	return &p
	//grepIn.Write([]byte("hello grep\ngoodbye grep"))
	//grepIn.Close()
	//grepBytes, _ := ioutil.ReadAll(grepOut)
	//grepCmd.Wait()
}

var connList []net.Conn
var subprocList []*subProx

func writeMessage(c net.Conn, m string) {
	w := bufio.NewWriter(c)
	w.Write([]byte(m))
	//w.Write([]byte("\n"))
	w.Flush()
}

func broadcast(Q chan connection) {
	for {
		m := <-Q
		//log.Println(m)

		var mess svarmrgo.Message
		_ = json.Unmarshal([]byte(m.raw), &mess)
		handleMessage(mess)
		//for _, v := range messages {
		//	Q <- v
		//}
		inMessages++
		for _, c := range connList {
			if c != nil && c != m.port {
				go writeMessage(c, m.raw) //FIXME use proper output queues so we can drop misbehaving clients
				outMessages++
			}
		}

		for _, c := range subprocList {
			if c != nil && c != m.handle {
				go c.In.Write([]byte(m.raw)) //FIXME use proper output queues so we can drop misbehaving clients
				outMessages++
			}
		}

	}
}

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	log.Println(m)
	switch m.Selector {
	//We don't need to reveal ourselves here, if we weren't running, the message couldn't get through
	//I guess we should for network situations
	//FIXME
	/*case "reveal-yourself":
	inQ <- connection{nil, nil, svarmrgo.WireFormat(svarmrgo.Message{Selector: "announce", Arg: "spine"})}
	*/
	case "start-module":
		go StartSubproc(m.Arg, []string{"pipes"})
		//go StartSubproc(fmt.Sprintf("%v.exe", m.Arg), []string{"pipes"})
	case "debug":
		log.Println(m.Arg)
	case "log":
		log.Println(m.Arg)
	case "error":
		log.Println(m.Arg)
	}
	return []svarmrgo.Message{}
}

//Read incoming messages and place them on the input queue
func handleConnection(conn net.Conn, Q chan connection) {
	//scanner := bufio.NewScanner(conn)
	reader := bufio.NewReader(conn)

	for {
		//for scanner.Scan() {
		t, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println("Client disconnected: ", err)
			return
		}
		var m connection = connection{conn, nil, t}
		Q <- m
		//time.Sleep(time.Millisecond * 200)
		//}
	}

}

var inQ chan connection

/*
func start_network() {

	ln, err := net.Listen("tcp", "0.0.0.0:4816")
	if err != nil {
		fmt.Printf("Couldn't open port 4816")
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		//fmt.Println("Client connected")
		if err != nil {
			// handle error
		}
		connList = append(connList, conn)
		go handleConnection(conn, inQ)
	}
}
*/

func main() {
	inQ = make(chan connection, 200)
	connList = make([]net.Conn, 0)
	//Don't run network sockets from the server anymore, run the relay module
	//to handle TCP socket clients
	//go start_network()
	go broadcast(inQ)

	//go StartSubproc("svarmr/clock.exe", []string{"pipes"})
	//go StartSubproc("gui/gui.exe", []string{"pipes"})
	for _, v := range os.Args[1:] {
		log.Println("Starting ", v)
		StartSubproc(v, []string{"pipes"})
	}
	go func() {
		//time.Sleep(5.0 * time.Second)
		inQ <- connection{nil, nil, svarmrgo.WireFormat(svarmrgo.Message{Selector: "user-notify", Arg: "Server started"})}
	}()
	for {
		time.Sleep(1.0 * time.Second)
	}
}
