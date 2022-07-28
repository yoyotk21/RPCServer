package server

// Object that will greet client
type Greeter struct {
	Message string
}

// Greets a person with a name
func (g *Greeter) SayHi(name string, reply *string) error {
	*reply = g.Message + " " + name + "!"
	return nil
}

// Makes a new greeter
func NewGreeter(message string) *Greeter {
	g := new(Greeter)
	g.Message = message
	return g
}

// Returns the message of the greeter
func (g *Greeter) GetMessage() string {
	return g.Message
}
