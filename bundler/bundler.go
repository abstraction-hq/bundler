package bundler

import (
	"sync"

	log "github.com/inconshreveable/log15"
)

var logger = log.New("module", "bundler")

type Bundler struct {
	lock sync.RWMutex
	stop chan struct{}
}

func NewBundler() (*Bundler, error) {
	return &Bundler{}, nil
}

func (b *Bundler) Start() error {
	logger.Info("Start bundler...")
	b.lock.Lock()
	defer b.lock.Unlock()

	// go b.RunCrawler()

	b.stop = make(chan struct{})
	return nil
}

func (b *Bundler) Wait() {
	b.lock.RLock()
	stop := b.stop

	b.lock.RUnlock()

	<-stop
}

func (b *Bundler) Stop() {
	logger.Warn("Stop bundler...")
	b.lock.Lock()
	defer b.lock.Unlock()

	close(b.stop)
}