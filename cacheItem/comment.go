package cacheItem

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

type String string

func (s String)Len()int{
	return len(s)
}