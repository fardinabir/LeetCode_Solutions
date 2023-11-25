
type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	valMap     map[int]*Node
	head, tail *Node
	cap, len   int
}

func Constructor(capacity int) LRUCache {
	head := &Node{-1, -1, nil, nil}
	tail := &Node{-2, -1, head, nil}
	head.next = tail

	return LRUCache{
		valMap: make(map[int]*Node, capacity),
		head:   head,
		tail:   tail,
		cap:    capacity,
		len:    0,
	}
}

func (cache *LRUCache) Get(key int) int {
	cur, ok := cache.valMap[key]
	if !ok {
		return -1
	}

	// key is already the head
	if cache.head.next == cur {
		return cur.value
	}

	// re-order otherwise
	cur.prev.next = cur.next
	cur.next.prev = cur.prev

	cur.next = cache.head.next
	cur.prev = cache.head
	cache.head.next.prev = cur
	cache.head.next = cur
	return cur.value

}

func (cache *LRUCache) Put(key int, value int) {
	if cache.Get(key) != -1 { // the key is available
		cache.head.next.value = value
		return
	}

	var cur *Node
	if cache.len >= cache.cap {
		cache.len--
		cur = cache.tail.prev
		cur.prev.next = cache.tail
		cache.tail.prev = cur.prev
		delete(cache.valMap, cur.key)
	} else {
		cur = &Node{}
	}

	cache.len++
	cur.value = value
	cur.key = key
	cur.next = cache.head.next
	cur.prev = cache.head
	cache.head.next.prev = cur
	cache.head.next = cur
	cache.valMap[key] = cur
}