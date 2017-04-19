//The message hub.  Server relays messages between programs
//
//Start with
//
//    server localhost 4816
package main
import (
    "log"
    "net"
    "bufio"
    "fmt"
	"os"

)

const queueLength int = 1

var inMessages int = 0
var outMessages int = 0

type connection struct {
    port net.Conn
    raw string
}

var connList []net.Conn

func writeMessage (c net.Conn, m string) {
                        w := bufio.NewWriter(c)
                        w.Write([]byte(m))
                        //w.Write([]byte("\n"))
                        w.Flush()
                        log.Println("Message sent, buffer at ", 
                    }

func broadcast(Q chan connection) {
    for {
            m := <- Q
            inMessages++
            for _, c := range outQs {
                //FIXME don't send back to the originator
                //if ( c.raw != nil && c.port != m.port) {
                    c <- m
                //}
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
                        //fmt.Println("Client disconnected: ", err)
                        return
				}
                var m connection = connection{ conn, t }
                Q <- m
				//time.Sleep(time.Millisecond * 200)
        //}
    }

}

var outQs []chan connection


func outWorker(c net.Conn, outQ chan connection) {
    for {
            m := <- outQ
            outMessages++
                if ( c != nil && c != m.port) {
                    go writeMessage(c,m.raw) //FIXME use proper output queues so we can drop misbehaving clients
                    outMessages++
            }
        }
}



func main() {
    log.Println("Initialising...")
    connList = make([]net.Conn,0)
    inQ := make(chan connection, queueLength)
    go broadcast(inQ)
    log.Println("Input queue started")
    ln, err := net.Listen("tcp", "0.0.0.0:4816")
    if err != nil {
          fmt.Printf("Couldn't open port 4816")
		      os.Exit(1)
    }
    log.Println("Listening on 0.0.0.0:4816")
    for {
        conn, err := ln.Accept()
        //fmt.Println("Client connected")
        if err != nil {
            // handle error
        }
        connList = append(connList, conn)
        outQ := make(chan connection, queueLength)
        outQs = append(outQs, outQ)
        go handleConnection(conn, inQ)
        go outWorker(conn, outQ)
        log.Println("Accepted incoming connection")
    }
}
