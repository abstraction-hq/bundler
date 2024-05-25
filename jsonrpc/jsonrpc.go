package jsonrpc

import (
	"github.com/abstraction-hq/abstraction-wallet-node/config"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/inconshreveable/log15"
)

var logger = log.New("module", "jsonrpc")

type JsonRpc struct {
	conf config.Config
}

func (j *JsonRpc) dump() {
	privateKey, err := crypto.HexToECDSA(j.conf.PrivateKey)
	if err != nil {
		logger.Error("Fail to create ECDSA", "detail", err)
	}

	logger.Info("Private key", "Detail", privateKey)
}

func NewJsonRpc(conf config.Config) (*JsonRpc, error) {
	return &JsonRpc{
		conf,
	}, nil
}
