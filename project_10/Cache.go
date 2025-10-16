package main

import "fmt"


type Cache interface{
	Set(key string, value string)
	Get(Key string)string
}


// Cache memory defination

type Basecache struct {
	Storage map[string]string
}

func ( b *Basecache) Set(key string , value string){
	b.Storage[key] = value
}

func (b *Basecache) Get(key string)string{
	return b.Storage[key]
}

// in memmory cache

type InMemoryCache struct{
	Basecache
}

func NewInmemoryCache()*InMemoryCache{
	return &InMemoryCache{Basecache{Storage: make(map[string]string)}}
}

// redis mock implementation

type RedisMock struct{
	Basecache
}

func NewRedisMock()*RedisMock{
	return &RedisMock{Basecache{Storage: make(map[string]string)}}
}

func ( r *RedisMock) Set(key string, value string) {
	fmt.Println("[RedisMock] simulating network call ...")
	r.Storage[key]= value
}

func  main(){
	var cache Cache

	fmt.Println("=== In memory cache =====")
	cache = NewInmemoryCache()
	cache.Set("language", "go")
	fmt.Println("Stored in memory : ", cache.Get("language"))


	fmt.Println("\n === Redis Mock ====")
	cache = NewInmemoryCache()
	cache.Set("framework","Gin")
	fmt.Println("stored in redis mock ", cache.Get("framework"))
}