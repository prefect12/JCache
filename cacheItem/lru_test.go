package cacheItem

import "testing"

func TestGet(t *testing.T) {
	lru := NewLruCache(int64(1), nil)
	lru.Add("key1", String("1,2,3"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1,2,3" {
		t.Fatalf("cache fail")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := NewLruCache(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}
