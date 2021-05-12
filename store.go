package lark

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrStoreNotFound = errors.New("store not found")

type Store interface {
	Get(ctx context.Context, key string) (string, time.Duration, error)
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

type StoreMemory struct {
	data map[string]*storeMemoryElem
	lock sync.Mutex
}

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

func (r *StoreMemory) Get(ctx context.Context, key string) (string, time.Duration, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	v, _ := r.data[key]
	if v == nil {
		return "", 0, ErrStoreNotFound
	}

	ttl := v.Expired.Sub(time.Now())
	if ttl >= time.Second*5 {
		return v.Data, ttl, nil
	}

	delete(r.data, key)

	return "", 0, ErrStoreNotFound
}

func (r *StoreMemory) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.data[key] = &storeMemoryElem{
		Data:    val,
		Expired: time.Now().Add(ttl),
	}

	return nil
}
