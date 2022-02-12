package JCache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Error("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGet(t *testing.T) {
	loadCount := make(map[string]int, len(db))
	myCache := NewGroup("score", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCount[key]; !ok {
					loadCount[key] = 1
				}
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	for k, v := range db {
		if view, err := myCache.Get(k); err != nil || view.String() != v {
			t.Fatalf("faild to get value from Tom")
		}
		if _, err := myCache.Get(k); err != nil || loadCount[k] > 1 {
			t.Fatalf("cache %s miss", k)
		} else {
			fmt.Println(loadCount)
		}
	}
	if view, err := myCache.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty,but %s got", view)
	}
}
