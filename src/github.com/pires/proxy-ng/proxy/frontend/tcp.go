package frontend

import (
	"bufio"
	"fmt"
	"net"
)

// package initializer
func init() {
	RegisterFactory("tcp", factory)
}

type TCP struct {
	channel chan net.Conn
	close   <-chan bool
}

func factory(close <-chan bool) Frontend {
	return &TCP{
		channel: make(chan net.Conn, 16),
		close:   close,
	}
}

func process(conn net.Conn) error {
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Received message: %s", msg)

	return nil
}

func (this *TCP) Process() {
	for {
		select {
		case conn := <-this.channel:
			process(conn)
		case <-this.close:
			return
		}
	}
}

func (this *TCP) GetConnectionChannel() chan<- net.Conn {
	return this.channel
}
