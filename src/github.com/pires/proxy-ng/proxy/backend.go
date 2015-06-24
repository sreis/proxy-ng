package proxy

type Backend interface {
	Reload() error
}
