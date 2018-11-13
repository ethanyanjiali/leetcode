package main

import (
	"fmt"
)

type ExamRoom struct {
}

func Constructor(N int) ExamRoom {
	return ExamRoom{}
}

func (this *ExamRoom) Seat() int {
	return 0
}

func (this *ExamRoom) Leave(p int) {
}

func main() {
	room := Constructor(10)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	room.Leave(4)
	fmt.Println(room.Seat())

	room = Constructor(10)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	room.Leave(0)
	room.Leave(4)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	room.Leave(0)
	room.Leave(4)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	room.Leave(7)
	fmt.Println(room.Seat())
	room.Leave(3)
	fmt.Println(room.Seat())
	room.Leave(3)
	fmt.Println(room.Seat())
	room.Leave(9)
	fmt.Println(room.Seat())
	room.Leave(0)
	room.Leave(8)
	fmt.Println(room.Seat())
	// room.Leave(0)
	// room.Leave(8)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
