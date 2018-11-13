package main

import "container/list"

type RecentCounter struct {
	Queue *list.List
	Limit int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Queue: list.New(),
		Limit: 3000,
	}
}

func (this *RecentCounter) Ping(t int) int {
	threshold := t - this.Limit
	length := this.Queue.Len()
	for i := 0; i < length; i++ {
		el := this.Queue.Back()
		if el.Value.(int) >= threshold {
			break
		}
		this.Queue.Remove(el)
	}
	this.Queue.PushFront(t)
	return this.Queue.Len()
}

func main() {
	counter := Constructor()
	counter.Ping(642)
	counter.Ping(1849)
	counter.Ping(4921)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
