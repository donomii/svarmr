package main
import (
    "net"
    "bufio"
    "fmt"
	"os"

)

var inMessages int = 0
var outMessages int = 0

type connection struct {
    port net.Conn
    raw string
}

var connList []net.Conn

func writeMessage (c connection) {
                        w := bufio.NewWriter(c)
                        w.Write([]byte(m.raw))
                        w.Write([]byte("\n"))
                        w.Flush()
                    }

func broadcast(Q chan connection) {
    for {
            m := <- Q
            inMessages++
            for _, c := range connList {
                if ( c != nil && c != m.port) {
                    go writeMessage(c) //FIXME use proper output queues so we can drop misbehaving clients
                        outMessages++
                }
            }
        }
}


//Read incoming messages and place them on the input queue
func handleConnection (conn net.Conn, Q chan connection) {
    //scanner := bufio.NewScanner(conn)
	reader := bufio.NewReader(conn)

    for {
            //for scanner.Scan() {
				t,err := reader.ReadString('\n')
				if err != nil {
                        fmt.Println("Client disconnected: ", err)
                        return
				}
                var m connection = connection{ conn, t }
                Q <- m
				//time.Sleep(time.Millisecond * 200)
        //}
    }

}

func main() {
    connList = make([]net.Conn,0)
    inQ := make(chan connection, 200)
    go broadcast(inQ)
    ln, err := net.Listen("tcp", "0.0.0.0:4816")
    if err != nil {
          fmt.Printf("Couldn't open port 4816")
		      os.Exit(1)
    }
    for {
        conn, err := ln.Accept()
        fmt.Println("Client connected")
        if err != nil {
            // handle error
        }
        connList = append(connList, conn)
        go handleConnection(conn, inQ)
    }
}
