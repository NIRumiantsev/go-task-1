package cache

type Data map[string]interface{}

type Cache struct {
	data Data
}

func (cache *Cache) Set(key string, value interface{}) {
	cache.data[key] = value
}

func (cache *Cache) Get(key string) interface{} {
	return cache.data[key]
}

func (cache *Cache) Delete(key string) {
	delete(cache.data, key)
}

func New() Cache {
	return Cache{data: make(Data)}
}
