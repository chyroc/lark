/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
