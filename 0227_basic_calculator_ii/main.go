package main

import "strconv"

func nextNum(expr string, start int) (int, int) {
	i := start
	for i < len(expr) && expr[i] >= '0' && expr[i] <= '9' {
		i += 1
	}
	v, _ := strconv.Atoi(string(expr[start:i]))
	return v, i
}

func calculate(s string) int {
	num, i := nextNum(s, 0)
	stack := []int{num}

	for i < len(s) {
		c := s[i]
		num, i = nextNum(s, i+1)
		if c == '+' {
			stack = append(stack, num)
		} else if c == '-' {
			stack = append(stack, -num)
		} else if c == '*' {
			last := stack[len(stack)-1]
			last = last * num
			stack[len(stack)-1] = last
		} else if c == '/' {
			last := stack[len(stack)-1]
			last = last / num
			stack[len(stack)-1] = last
		}
	}

	res := 0
	for _, num := range stack {
		res += num
	}

	return res
}

func main() {

}
