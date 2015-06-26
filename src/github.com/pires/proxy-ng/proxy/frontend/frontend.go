package frontend

import "net"

type Frontend interface {
	GetConnectionChannel() chan<- net.Conn
	Process()
}

type FrontendFactory func(close <-chan bool) Frontend

var frontends = map[string]FrontendFactory{}

//
func RegisterFactory(protocol string, factory FrontendFactory) error {
	// TODO: make frontend a type?
	// TODO: error handling
	frontends[protocol] = factory
	return nil
}

//
func GetFactory(frontend string) FrontendFactory {
	// TODO: error handling
	return frontends[frontend]
}
