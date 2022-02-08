package cacheItem

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}
