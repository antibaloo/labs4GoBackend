package main

import (
	"fmt"
	"sync"
	"time"
)

var _ Cache = InMemoryCache{} // это трюк для проверки типа: до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn  time.Duration
	dataMutex *sync.Mutex
	data      map[string]CacheEntry
}

func (i InMemoryCache) Set(key string, value interface{}) {
	i.dataMutex.Lock()
	i.data[key] = CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
	i.dataMutex.Unlock()
}
func (i InMemoryCache) Get(key string) interface{} {
	i.dataMutex.Lock()
	defer i.dataMutex.Unlock()
	value, ok := i.data[key]
	if !ok {
		return "there is no such key"
	}

	if i.expireIn < time.Since(i.data[key].settledAt) {
		delete(i.data, key)
		return "data you need id expired"
	} else {
		return value.value
	}

}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn:  expireIn,
		dataMutex: &sync.Mutex{},
		data:      make(map[string]CacheEntry),
	}
}

func main() {
	myCache := NewInMemoryCache(15 * time.Second)
	myCache.Set("Andrey", 15)
	myCache.Set("Anna", "wife")
	duration := 5 * time.Second
	time.Sleep(duration)
	fmt.Println(myCache.Get("Andrey"))
	fmt.Println(myCache.Get("Anna"))

	duration = 15 * time.Second
	time.Sleep(duration)
	fmt.Println(myCache.Get("Andrey"))
	fmt.Println(myCache.Get("Anna"))

	fmt.Println(myCache.Get("Andrey"))
	fmt.Println(myCache.Get("Anna"))
}
