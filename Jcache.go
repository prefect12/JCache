package JCache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string)([]byte,error)
}

type GetterFunc func(key string)([]byte,error)

func(f GetterFunc)Get(key string)([]byte,error){
	return f(key)
}

type Group struct {
	name string
	getter Getter
	mainCache cache
}

var(
	mu sync.RWMutex
	groups = make(map[string]*Group)
)

func newGroup(name string,cacheBytes int64,getter Getter)*Group{
	if getter == nil{
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:name,
		getter:getter,
		mainCache:cache{cacheBytes:cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string)*Group{
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group)Get(key string)(Byteview,error){
	if key == ""{
		return Byteview{},fmt.Errorf("key is required")
	}
	if v,ok := g.mainCache.get(key);ok{
		log.Println("[JCache] hit")
		return v,nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (value Byteview, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (Byteview, error) {
	bytes,err := g.getter.Get(key)
	if err != nil{
		return Byteview{},err
	}
	value := Byteview{b:cloneBytes(bytes)}
	g.populateCache(key,value)
	return value,nil
}

func (g *Group) populateCache(key string, value Byteview) {
	g.mainCache.add(key,value)
}