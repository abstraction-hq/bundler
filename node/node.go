package node

import (
	"sync"

	log "github.com/inconshreveable/log15"
)

var logger = log.New("module", "node")

type Node struct {
	lock sync.RWMutex
	stop chan struct{}
}

func NewNode() (*Node, error) {
	return &Node{}, nil
}

func (b *Node) Start() error {
	logger.Info("Start node...")
	b.lock.Lock()
	defer b.lock.Unlock()

	// TODO: add services

	b.stop = make(chan struct{})
	return nil
}

func (b *Node) Wait() {
	b.lock.RLock()
	stop := b.stop

	b.lock.RUnlock()

	<-stop
}

func (b *Node) Stop() {
	logger.Warn("Stop node...")
	b.lock.Lock()
	defer b.lock.Unlock()

	close(b.stop)
}