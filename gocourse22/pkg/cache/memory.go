package cache

import (
	"sync"
	"time"
)

// Item представляє одиницю даних у кеші.
type Item struct {
	Value      interface{}
	Expiration int64
}

// Cache структура для кешування даних.
type Cache struct {
	items sync.Map
}

// NewCache створює новий екземпляр кешу.
func NewCache() *Cache {
	return &Cache{}
}

// Set додає або оновлює значення в кеші з вказаним часом життя (TTL).
// Якщо ttl == 0, елемент не буде мати часу закінчення.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}
	c.items.Store(key, Item{
		Value:      value,
		Expiration: expiration,
	})
}

// Get отримує значення за ключем з кешу.
// Повертає значення та true, якщо елемент знайдено і не прострочено.
func (c *Cache) Get(key string) (interface{}, bool) {
	item, ok := c.items.Load(key)
	if !ok {
		return nil, false
	}
	cacheItem := item.(Item)
	if cacheItem.Expiration > 0 && time.Now().UnixNano() > cacheItem.Expiration {
		c.Delete(key) // Видаляємо прострочене значення
		return nil, false
	}
	return cacheItem.Value, true
}

// Delete видаляє елемент з кешу.
func (c *Cache) Delete(key string) {
	c.items.Delete(key)
}

// Size повертає кількість елементів в кеші.
func (c *Cache) Size() int {
	size := 0
	c.items.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
}
