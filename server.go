package main
import (
    "net"
    "bufio"
    "fmt"
	"time"
	
)

type connection struct {
    port net.Conn
    raw string
}

var connList []net.Conn

func broadcast(Q chan connection) {
    for {
		fmt.Printf("Outer broadcast loop")
            m := <- Q
            for _, c := range connList {
                if ( c != nil && c != m.port) {
                    c.Write([]byte(m.raw))
                    c.Write([]byte("\r\n"))
					fmt.Printf("Inner broadcast loop")
                }
            }
        }
}
func handleConnection (conn net.Conn, Q chan connection) {
    //scanner := bufio.NewScanner(conn)
	reader := bufio.NewReader(conn)

    for {
			fmt.Printf("Outer handler loop")
            //for scanner.Scan() {
				t,err := reader.ReadString('\n')
				fmt.Printf("%V\n", err)
				if err != nil {
					return
				}
                var m connection = connection{ conn, t }
                Q <- m
                fmt.Printf("Inner scanner loop")
				time.Sleep(time.Millisecond * 200)
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
		fmt.Printf("Listener loop")
        conn, err := ln.Accept()
        fmt.Println("Caught connection")
        if err != nil {
            // handle error
        }
        connList = append(connList, conn)
        go handleConnection(conn, inQ)
    }
}
