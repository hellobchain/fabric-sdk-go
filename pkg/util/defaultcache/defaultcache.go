package defaultcache

import (
	"github.com/wsw365904/fabric-sdk-go/pkg/util/cache"
	"sync"
)

var dc *defaultCache

type defaultCache struct {
	cache *cache.Cache
	lock  sync.Mutex
}

func newDefaultCache() *defaultCache {
	return &defaultCache{
		cache: cache.NewCache(),
	}
}

var initOnce sync.Once

func DefaultCache() *defaultCache {
	initOnce.Do(func() {
		dc = newDefaultCache()
	})
	return dc
}

func (d *defaultCache) Set(key string, value interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.cache.Del(key)
	d.cache.Set(key, value, -1)

}

func (d *defaultCache) Get(key string) (interface{}, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	get, err := d.cache.Get(key)
	if err != nil {
		return nil, err
	}
	return get, nil
}

func (d *defaultCache) Delete(key string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.cache.Del(key)
}
