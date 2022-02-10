package cacheItem

import (
	"fmt"
	"testing"
)



func TestLfuCache_GetGet(t *testing.T){
	lru := NewLfuCache(int64(1),nil)
	lru.Add("key1",String("1,2,3"))
	if v,ok := lru.Get("key1");!ok||string(v.(String)) != "1,2,3"{
		t.Fatalf("cache fail")
	}else{
		fmt.Println("get success value:",v)
	}

	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestLfuCache_Get(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := NewLfuCache(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	lru.Get(k3)
	lru.Get(k2)

	fmt.Println(lru.Get(k1))
	fmt.Println(lru.Get(k2))
	fmt.Println(lru.Get(k3))
	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}