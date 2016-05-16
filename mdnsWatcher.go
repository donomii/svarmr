package main
import (
    "fmt"
    //"os"
    //"encoding/json"
    "github.com/donomii/svarmrgo"
    "time"
"net"
)

import "github.com/hashicorp/mdns"

//type ServiceEntry struct {
//>···Name       string
//>···Host       string
//>···AddrV4     net.IP
//>···AddrV6     net.IP
//>···Port       int·
//>···Info       string
//>···InfoFields []string
//
//>···Addr net.IP // @Deprecated
//
//>···hasTXT bool
//>···sent   bool
//}

func watchDNS (entriesCh chan *mdns.ServiceEntry) {
    for {
        // Start the lookup
        go mdns.Lookup("_svarmr._tcp", entriesCh)
        time.Sleep(600000 * time.Millisecond)
    }
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
     switch m.Selector {
        case "reveal-yourself" :
           svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "mDnsWatcher"})
      }
}

func main() {
    conn := svarmrgo.CliConnect()
    go svarmrgo.HandleInputs(conn, handleMessage)
    entriesCh := make(chan *mdns.ServiceEntry, 4) 
    go watchDNS(entriesCh)
    go func() {
        for entry := range entriesCh {
           svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "mdns-found-ipv4", Arg: fmt.Sprintf("%v:%v", entry.AddrV4, entry.Port)})
           svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "mdns-found-ipv6", Arg: fmt.Sprintf("%v:%v", entry.AddrV6, entry.Port)})
        }
    }()
    for{
        time.Sleep(120000 * time.Millisecond)
    }
}
