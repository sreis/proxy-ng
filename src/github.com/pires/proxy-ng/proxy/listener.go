package proxy

type Listener interface {
	Listen() error
	Accept() error
	Close() error
	Reload() error
}
