package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash map from bytes to uint32
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int            // multiple of the dummy node
	keys     []int          // hash ring
	hashmap  map[int]string // dummy node & actual node relation
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashmap:  make(map[int]string),
	}
	if m.hash == nil { // if have not the hash func, use the defalut
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add keys to the hash
func (m *Map) Add(keys ...string) { // can give lots of keys
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ { // set the replica multiple dummy nodes
			dummyhash := int(m.hash([]byte(strconv.Itoa(i) + key))) // Itoa: int to stirng
			m.keys = append(m.keys, dummyhash)                      // add to hash ring
			m.hashmap[dummyhash] = key                              // add map relation
		}
	}
	sort.Ints(m.keys) // since we should search the key clockwise
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key))) // calculate hash for key
	// clockwise for searching
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashmap[m.keys[idx%len(m.keys)]] // since it is a ring e.g.
	// if we get the i is len(m.key), so the idx should be 0
}
