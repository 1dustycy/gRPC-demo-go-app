package multiservices

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
)

type registry struct {
	mods map[string]mod
	mu   sync.Mutex
}

type mod interface {
	Start() error
	Close() error
}

var (
	defaultRegistry *registry
	healthy         int32
)

// Init ...
func Init() {
	defaultRegistry = &registry{
		mods: make(map[string]mod),
		mu:   sync.Mutex{},
	}
	atomic.StoreInt32(&healthy, 1)
}

// Healthz ...
func Healthz() bool {
	if ok := atomic.LoadInt32(&healthy); ok != 1 {
		return false
	}

	return true
}

// AddMod ...
func AddMod(id string, m mod) error {
	if defaultRegistry == nil {
		return errors.New("not initialized")
	}

	if h := atomic.LoadInt32(&healthy); h != 1 {
		return errors.New("unhealthy service")
	}

	if m == nil {
		return errors.New("mod is nil")
	}

	defaultRegistry.mu.Lock()
	defer defaultRegistry.mu.Unlock()

	if _, exists := defaultRegistry.mods[id]; exists {
		return errors.New("mod ID already exists")
	}

	defaultRegistry.mods[id] = m

	go func() {
		if err := m.Start(); err != nil {
			atomic.StoreInt32(&healthy, 0)
		}
	}()

	log.Printf("mod %s is registered and running", id)

	return nil
}

// Shutdown ..
func Shutdown() {
	defaultRegistry.mu.Lock()
	defer defaultRegistry.mu.Unlock()

	atomic.StoreInt32(&healthy, 0)

	for id := range defaultRegistry.mods {
		if err := defaultRegistry.mods[id].Close(); err != nil {
			log.Printf("mod %s was not properly closed: %v", id, err)
		}
	}

	log.Println("[multimods] successful shutdown")
}
