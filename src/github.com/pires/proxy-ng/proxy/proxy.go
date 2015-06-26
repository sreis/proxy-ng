package proxy

import (
	"github.com/pires/proxy-ng/proxy/frontend"
	"github.com/pires/proxy-ng/proxy/listener"
)

type Proxy struct {
	close     <-chan bool
	listeners map[int]*listener.Listener // map port with attached listener
}

func New(close <-chan bool) *Proxy {
	return &Proxy{
		close:     close,
		listeners: make(map[int]*listener.Listener),
	}
}

// spawns go routine that is listening on a socket at specified port
func (this *Proxy) NewListener(protocol string, address string, port int) error {

	// create a TCP frontend that is listening on port 8080
	constructor := frontend.GetFactory(protocol)
	f := constructor(this.close)

	l := listener.New(address, port, this.close, f.GetConnectionChannel())
	go l.Listen()

	this.listeners[port] = l

	return nil
}
