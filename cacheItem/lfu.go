package cacheItem

import (
	"container/heap"
)

type Frequency []*freqItem

type freqItem struct {
	Entry *entry
	frequency int
}



func (f Frequency)Len()int{
	return len(f)
}

func(f Frequency)Less(i,j int)bool{
	return f[i].frequency < f[j].frequency
}

func(f Frequency)Swap(i,j int){
	f[i],f[j] = f[j],f[i]
}

func(f *Frequency)Push(x interface{}){
	item := x.(*freqItem)
	*f = append(*f,item)
}

func(f *Frequency)Pop()interface{}{
	old := *f
	item := old[len(*f)-1]
	*f = old[:len(*f)-1]
	return item
}


type LfuCache struct {
	maxBytes int64
	nbytes int64
	cache map[string]*freqItem
	frequency *Frequency
	OnEvicted func(key string,value Value)
}

func NewLfuCache(maxBytes int64,onEvicate func(string, Value)) *LfuCache {
	myFrequency := &Frequency{}
	heap.Init(myFrequency)
	return &LfuCache{
		maxBytes: maxBytes,
		cache: make(map[string]*freqItem),
		OnEvicted: onEvicate,
		frequency:myFrequency,
	}
}

func (c *LfuCache)Get(key string)(value Value,ok bool){
	if ele,ok := c.cache[key];ok{
		kv := ele.Entry
		ele.frequency += 1
		return kv.value,true
	}
	return nil,false
}

func (c *LfuCache)RemoveOldest(){
	item := c.frequency.Pop().(*freqItem)
	delete(c.cache,item.Entry.key)
	c.nbytes -= int64(len(item.Entry.key)) + int64(item.Entry.value.Len())
	if c.OnEvicted != nil{
		c.OnEvicted(item.Entry.key,item.Entry.value)
	}
}

func(c *LfuCache)Add(key string,value Value){
	if ele,ok := c.cache[key];ok{
		item := ele.Entry
		c.nbytes += int64(value.Len()) - int64(item.value.Len())
		ele.frequency += 1
	}else{
		ele := &freqItem{
			Entry: &entry{
				key:key,
				value: value,
			},
			frequency: 1,
		}
		heap.Push(c.frequency,ele)
		c.cache[key] = ele
		c.nbytes += int64(value.Len()) + int64(len(key))
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes{
		c.RemoveOldest()
	}
}

func (c *LfuCache)Len()int{
	return len(c.cache)
}



