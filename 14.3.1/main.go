package main

import "fmt"

type hashmap struct {
	data map[uint64]string
}

// Хэш-функция из 14.2.2
func (h *hashmap) hashstr(val string) uint64 {
	var sum int
	for i, s := range val {
		sum += (i + 1) * int(s)
	}
	return uint64(sum)
}

// Запись элемента
func (h *hashmap) Set(key, val string) {
	h.data[h.hashstr(key)] = val
}
func NewHashMap() *hashmap {
	return &hashmap{data: make(map[uint64]string)}
}

// Получение элемента
func (h *hashmap) Get(key string) (value string, ok bool) {
	if value, ok := h.data[h.hashstr(key)]; ok {
		return value, true
	} else {
		return "", false
	}
}

// Удаление элемента
func (h *hashmap) Delete(key string) {
	delete(h.data, h.hashstr(key))
}

func main() {
	var hashMap = NewHashMap()
	var requests = []string{"abalakov", "abalakova", "ivanov"}
	hashMap.Set("abalakov", "abalakov andrey sergeevich")
	hashMap.Set("abalakova", "abalakova anna vladimirovna")

	fmt.Println("getting data for keys:", requests)
	for _, key := range requests {
		if value, ok := hashMap.Get(key); ok {
			fmt.Println(key, ":", value)
		} else {
			fmt.Println("no", key)
		}
	}

	fmt.Println("deleting ivanov and abalakov")

	hashMap.Delete("ivanov")
	hashMap.Delete("abalakov")

	fmt.Println("getting data for keys:", requests)

	for _, key := range requests {
		if value, ok := hashMap.Get(key); ok {
			fmt.Println(key, ":", value)
		} else {
			fmt.Println("no", key)
		}
	}

}
