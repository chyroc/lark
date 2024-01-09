package lark_ws

import (
	"sync"
	"time"
)

type cache struct {
	values sync.Map
}

func newCache() *cache {
	return &cache{
		values: sync.Map{},
	}
}

func (r *cache) get(key string) (interface{}, bool) {
	val, ok := r.values.Load(key)
	if !ok {
		return nil, false
	}
	item := val.(*cacheItem)
	if item.expired.Before(time.Now()) {
		r.values.Delete(key)
		return nil, false
	}
	return item.val, true
}

func (r *cache) set(key string, val interface{}) {
	r.values.Store(key, &cacheItem{
		val: val,

		// https://open.feishu.cn/document/server-docs/event-subscription-guide/overview
		expired: time.Now().Add(time.Hour),
	})
}

type cacheItem struct {
	val     interface{}
	expired time.Time
}
