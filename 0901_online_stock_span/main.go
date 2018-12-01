package main

import "fmt"

type StockSpanner struct {
    stack [][2]int 
}


func Constructor() StockSpanner {
    return StockSpanner{
    	stack: [][2]int{},
    }
}


func (this *StockSpanner) Next(price int) int {
	days := 1
	// keep popping the stack when the top element is smaller than current price
	for len(this.stack) > 0 && this.stack[len(this.stack) - 1][0] <= price {
		top := this.stack[len(this.stack) - 1]
		// since the top element has smaller price, everthing smaller than this element should be smaller than current price
		// so we add them on
		days += top[1]
		this.stack = this.stack[:len(this.stack) - 1]
	}
	// aftering poping smaller elements, the top of stack should be bigger than current price
	// so we put this new price on the top
	this.stack = append(this.stack, [2]int{price, days})
	return days
}

func main() {
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	s := Constructor()
	for _, price := range prices {
		fmt.Println(s.Next(price))
	}
}

