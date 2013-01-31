package template

import (
	"sync"
)

type nameSpace struct {
	mu    sync.RWMutex
	tmpls map[string]*Template
}

func (n *nameSpace) get(name string) (*Template, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	t, b := n.tmpls[name]
	return t, b
}

func (n *nameSpace) Len() int {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return len(n.tmpls)
}

func (n *nameSpace) Get(name string) *Template {
	t, _ := n.get(name)
	return t
}

func (n *nameSpace) Add(name string, t *Template) {
	if _, ok := n.get(name); !ok {
		n.mu.Lock()
		n.tmpls[name] = t
		n.mu.Unlock()
	}
}

func (n *nameSpace) Del(name string) {
	if _, ok := n.get(name); ok {
		n.mu.Lock()
		n.mu.Unlock()
		delete(n.tmpls, name)
	}
}

func (n *nameSpace) Set(name string, t *Template) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.tmpls[name] = t
}
