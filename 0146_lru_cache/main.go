package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Key   int
	Value int
}

type LRUCache struct {
	HashMap  map[int]*list.Element
	Queue    *list.List
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		map[int]*list.Element{},
		&list.List{},
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if el, ok := this.HashMap[key]; ok {
		// update the recency
		this.Queue.MoveToFront(el)
		return el.Value.(Node).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if el, ok := this.HashMap[key]; ok {
		// update the record with new value
		el.Value = Node{key, value}
		this.Queue.MoveToFront(el)
	} else {
		if len(this.HashMap) >= this.Capacity {
			// remove last one
			last := this.Queue.Back()
			delete(this.HashMap, last.Value.(Node).Key)
			this.Queue.Remove(last)
		}
		// prepend to the front
		el = this.Queue.PushFront(Node{key, value})
		this.HashMap[key] = el
	}
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
