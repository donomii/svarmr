package main
import (
"errors"
"log"
    "fmt"
    //"os"
    //"encoding/json"
    "github.com/donomii/svarmrgo"
    "time"
"net"
"math/rand"
 "github.com/oleksandr/bonjour"
 "os"
)



func random(min, max int) int {
    return rand.Intn(max - min) + min
}

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



func watchDNS (conn net.Conn, entriesCh chan *bonjour.ServiceEntry) {
svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "debug", Arg: "Scanning mdns for _tcp"})
resolver, err := bonjour.NewResolver(nil)
    if err != nil {
        log.Println("Failed to initialize resolver:", err.Error())
        os.Exit(1)
    }
		log.Println("Starting mdns lookup")
		
		//, "_smb._tcp", "_googlecast._tcp", "_udisks-ssh._tcp",   "_webdav._tcp", "_ssh._tcp", "_telnet._tcp"
			go func (){
				err = resolver.Browse("_tcp", "", entriesCh)
				if err != nil {
					log.Println("Failed to browse:", err.Error())
				}
			}()
			//log.Println("Searching for ", v)
		
        
		//"_svarmr._tcp.",
		
		
		
		time.Sleep(time.Duration(random(50000,300000))*time.Millisecond)
    
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

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	out := []svarmrgo.Message{}
     switch m.Selector {
        case "reveal-yourself" :
           m.Respond(svarmrgo.Message{Selector: "announce", Arg: "mDnsWatcher"})
      }
	  return out
}

func main() {
    rand.Seed(time.Now().Unix())
	conn := svarmrgo.CliConnect()
	go svarmrgo.HandleInputLoop(conn, handleMessage)
    entriesCh := make(chan *bonjour.ServiceEntry, 40) 
    go watchDNS(conn, entriesCh)
    go func() {
        for entry := range entriesCh {
			log.Println("Got mdns response")
           //svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "mdns-found-svarmr-ipv4", Arg: fmt.Sprintf("%v:%v", entry.AddrIPv4, entry.Port)})
           //svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "mdns-found-svarmr-ipv6", Arg: fmt.Sprintf("%v:%v", entry.AddrIPv6, entry.Port)})
		   svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "mdns-service-found-summary", Arg: fmt.Sprintf("%v:%v (%v)", entry.HostName, entry.Port, entry.Instance)})
		   svarmrgo.SendMessage(conn, svarmrgo.Message{Selector: "mdns-service-found", Args: []string{
		   "Instance", fmt.Sprintf("%v", entry.Instance),
		   "HostName", fmt.Sprintf("%v", entry.HostName), 
		   "Text", fmt.Sprintf("%v", entry.Text), 
		   "AddrV4", fmt.Sprintf("%v", entry.AddrIPv4), 
		   "AddrV6", fmt.Sprintf("%v", entry.AddrIPv6), 
		   "Port", fmt.Sprintf("%v", entry.Port),
		   "Service", fmt.Sprintf("%v", entry.Service), 
		   "Domain", fmt.Sprintf("%v", entry.Domain), 
		   "TTL", fmt.Sprintf("%v", entry.TTL), 
		   }})
        }
    }()
    for{
        time.Sleep(12000 * time.Millisecond)
    }
}
