package main

import (
	"fmt"
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
	expireIn time.Duration
	data     map[string]CacheEntry
}

func (i InMemoryCache) Set(key string, value interface{}) {
	i.data[key] = CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
}
func (i InMemoryCache) Get(key string) interface{} {
	if value, ok := i.data[key]; ok {
		if i.expireIn < time.Since(i.data[key].settledAt) {
			delete(i.data, key)
			return "data you need id expired"
		} else {
			return value.value
		}
	} else {
		return "there is no such key"
	}

}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,
		data:     make(map[string]CacheEntry),
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
