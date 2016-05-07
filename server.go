package main
import (
    "net"
    "bufio"
    "fmt"
)

type message struct {
    port net.Conn
    raw string
}

var connList []net.Conn

func broadcast(Q chan message) {
    for {
            m := <- Q
            for _, c := range connList {
                if ( c != nil && c != m.port) {
                    c.Write([]byte(m.raw))
                    c.Write([]byte("\r\n"))
                }
            }
        }
}
func handleConnection (conn net.Conn, Q chan message) {
    scanner := bufio.NewScanner(conn)

    for {
            for scanner.Scan() {
                var m message = message{ conn, scanner.Text() }
                Q <- m
        }
    }

}

func main() {
    connList = make([]net.Conn,0)
    inQ := make(chan message, 200)
    go broadcast(inQ)
    ln, err := net.Listen("tcp", "0.0.0.0:4816")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept()
        fmt.Println("Caught connection")
        if err != nil {
            // handle error
        }
        connList = append(connList, conn)
        go handleConnection(conn, inQ)
    }
}
