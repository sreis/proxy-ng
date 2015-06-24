package proxy

// Where magic happens

type Proxy struct {
	listeners map[Listener]string
}

func (p *Proxy) AddListener(listener *Listener) error {
	// TODO
	return nil
}

func (p *Proxy) RemoveListener(listener *Listener) error {
	// TODO
	return nil
}
