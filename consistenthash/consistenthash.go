package consistenthash

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int //sorted
	hashMap  map[int]string
}
