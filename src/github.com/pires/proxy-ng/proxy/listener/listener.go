package listener

import (
	"fmt"
	"net"
)

type Listener struct {
	address string
	close   <-chan bool
	out     chan<- net.Conn
}

func New(address string, port int, close <-chan bool, out chan<- net.Conn) *Listener {
	return &Listener{
		address: fmt.Sprintf("%s:%d", address, port),
		close:   close,
		out:     out,
	}
}

// Creates and binds socket to address, when a connection arrives
// forwards it to out channel.
func (this *Listener) Listen() error {

	ln, err := net.Listen("tcp", this.address)
	defer ln.Close()

	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// TODO: handle error
		}

		this.out <- conn
	}
}
