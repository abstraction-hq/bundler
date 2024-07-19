package jsonrpc

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/abstraction-hq/bundler/config"
	log "github.com/inconshreveable/log15"
)

var logger = log.New("module", "jsonrpc")

type JsonRpc struct {
	conf *config.Config
}

type TimeServer int64

func NewJsonRpc(conf *config.Config) (*JsonRpc, error) {
	return &JsonRpc{
		conf,
	}, nil
}

func (j *JsonRpc) Start() error {
	timeserver := new(TimeServer)
	// Register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// Listen for requests on port 1234
	l, e := net.Listen("tcp", ":2233")
	if e != nil {
		logger.Error("listen error:", "error", e)
	}
	http.Serve(l, nil)

	logger.Info("JsonRPC server started")
	return nil
}
