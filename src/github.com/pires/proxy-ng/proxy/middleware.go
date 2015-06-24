package proxy

type Middleware interface {
	Register() error
	Reload() error
}
