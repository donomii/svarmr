package chat

type Message struct {
	Selector    string `json:"Selector"`
	Arg         string `json:"Arg"`
	NamedArgs   map[string]string `json:"NamedArgs"`
    
}

func (self *Message) String() string {
	return "stringify: " + self.Selector + " : " + self.Arg
}
