package lark

import (
	"context"
	"errors"
	"sync"
	"time"
)

// ErrStoreNotFound ...
var ErrStoreNotFound = errors.New("store not found")

// Store ...
type Store interface {
	Get(ctx context.Context, key string) (string, time.Duration, error)
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

// StoreMemory ...
type StoreMemory struct {
	data map[string]*storeMemoryElem
	lock sync.Mutex
}

// NewStoreMemory ...
func NewStoreMemory() Store {
	return &StoreMemory{
		data: map[string]*storeMemoryElem{},
		lock: sync.Mutex{},
	}
}

type storeMemoryElem struct {
	Data    string
	Expired time.Time
}

// Get ...
func (r *StoreMemory) Get(ctx context.Context, key string) (string, time.Duration, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	v, _ := r.data[key]
	if v == nil {
		return "", 0, ErrStoreNotFound
	}

	ttl := v.Expired.Sub(timeNow())
	if ttl >= time.Second*5 {
		return v.Data, ttl, nil
	}

	delete(r.data, key)

	return "", 0, ErrStoreNotFound
}

// Set ...
func (r *StoreMemory) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.data[key] = &storeMemoryElem{
		Data:    val,
		Expired: timeNow().Add(ttl),
	}

	return nil
}

var timeNow = time.Now
