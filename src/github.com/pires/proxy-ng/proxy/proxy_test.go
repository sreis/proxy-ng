package proxy

import (
	"testing"
)

func TestNewTCPListener(t *testing.T) {
	protocol := "tcp"
	address := "localhost"
	port := 8080

	close := make(chan bool)

	proxy := New(close)
	proxy.NewListener(protocol, address, port)
}
