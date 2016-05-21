package main
import (
"errors"
    "fmt"
    //"os"
    //"encoding/json"
    "github.com/donomii/svarmrgo"
    "time"
"net"
"os"
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

var server *mdns.Server

func watchDNS (entriesCh chan *mdns.ServiceEntry) {
    for {
        // Start the lookup
        go mdns.Lookup("_svarmr._tcp.", entriesCh)
        time.Sleep(6000 * time.Millisecond)
    }
}

//https://code.google.com/archive/p/whispering-gophers/
func externalIP() (net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return nil, err
        }
        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue // not an ipv4 address
            }
            return ip, nil
        }
    }
    return nil, errors.New("are you connected to the network?")
}

func advertiseDNS () {
    // Setup our service export
    host, _ := os.Hostname()
    info := []string{"Svarmr network control bus"}
    ip, _ := externalIP()
    fmt.Printf("External IP: %v\n", ip)
    service, err := mdns.NewMDNSService(host, "_svarmr._tcp.", "", fmt.Sprintf("%v.local.", host), 4816, []net.IP{ip}, info)
    fmt.Printf("Error: %v\n", err)

    // Create the mDNS server, defer shutdown
    server, _ = mdns.NewServer(&mdns.Config{Zone: service})
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
           svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "mdns-found-svarmr-ipv4", Arg: fmt.Sprintf("%v:%v", entry.AddrV4, entry.Port)})
           svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "mdns-found-svarmr-ipv6", Arg: fmt.Sprintf("%v:%v", entry.AddrV6, entry.Port)})
        }
    }()
    advertiseDNS()
    for{
        time.Sleep(12000 * time.Millisecond)
    }
}
