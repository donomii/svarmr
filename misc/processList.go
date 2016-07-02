package main

import (
    "encoding/json"
    "regexp"
    "strings"
    "net"
    "os"
    "os/exec"
    "bytes"
    "github.com/donomii/svarmrgo"
    "fmt"
)

var history [3]string
var hpointer int
var lastNote string


func quickCommand (cmd *exec.Cmd) string{
    fmt.Println()
    fmt.Println("Started command")
    in := strings.NewReader("")
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stderr = &err
	cmd.Run()
    //fmt.Printf("Command result: %v\n", res)
	ret := fmt.Sprintf("%s", out)
    //fmt.Println(ret)
    return ret
}

func RegSplit(text string, delimeter string) []string {
    reg := regexp.MustCompile(delimeter)
    indexes := reg.FindAllStringIndex(text, -1)
    laststart := 0
    result := make([]string, len(indexes) + 1)
    for i, element := range indexes {
            result[i] = text[laststart:element[0]]
            laststart = element[1]
    }
    result[len(indexes)] = text[laststart:len(text)]
    return result
}

func getPS() []map[string]string {
 cmd := exec.Command("/bin/ps", "auxwww")
    res := quickCommand(cmd)
    lines := strings.Split(res, "\n")

    out := []map[string]string{}
    for _,l := range lines {
        b := RegSplit(l, "\\s+")
        if len(b) > 9 {
            r := map[string]string{}
            r["id"] = b[1]
            r["cpu"] = b[2]
            r["mem"] = b[3]
            r["time"] = b[9]
            r["command"] = b[10]
            out = append(out, r)
        }
    }
    //fmt.Printf("%v\n", out)
    return out
}

func handleMessage (conn net.Conn, m svarmrgo.Message) {
    switch m.Selector {
         case "reveal-yourself" :
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "announce", Arg: "noteProcessor"})
            case "shutdown" :
            os.Exit(0)
         case "get-process-list" :
            b, err := json.Marshal(getPS())
            if err != nil {
                fmt.Println("error:", err)
                os.Exit(1)
            }
            svarmrgo.RespondWith(conn, svarmrgo.Message{Selector: "process-list", Arg: string(b)})
    }
}

func main() {
    getPS()
    conn := svarmrgo.CliConnect()
    svarmrgo.HandleInputs(conn, handleMessage)
}
