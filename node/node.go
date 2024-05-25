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

func (n *Node) Start() error {
	logger.Info("Start node...")
	n.lock.Lock()
	defer n.lock.Unlock()

	// TODO: add services

	n.stop = make(chan struct{})
	return nil
}

func (n *Node) Wait() {
	n.lock.RLock()
	stop := n.stop

	n.lock.RUnlock()

	<-stop
}

func (n *Node) Stop() {
	logger.Warn("Stop node...")
	n.lock.Lock()
	defer n.lock.Unlock()

	close(n.stop)
}