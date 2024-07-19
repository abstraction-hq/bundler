package node

import (
	"sync"

	"github.com/abstraction-hq/bundler/config"
	"github.com/abstraction-hq/bundler/jsonrpc"
	log "github.com/inconshreveable/log15"
)

var logger = log.New("module", "node")

type Node struct {
	conf *config.Config
	lock sync.RWMutex
	rpc  *jsonrpc.JsonRpc
	stop chan struct{}
}

func NewNode(conf *config.Config) (*Node, error) {
	return &Node{
		conf: conf,
	}, nil
}

func (n *Node) Start() error {
	logger.Info("Start node...")
	n.lock.Lock()
	defer n.lock.Unlock()

	// TODO: add services
	n.rpc, _ = jsonrpc.NewJsonRpc(n.conf)
	n.rpc.Start()

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
