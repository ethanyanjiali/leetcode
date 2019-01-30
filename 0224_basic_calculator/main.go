package main

import "strconv"
import "strings"

func nextNum(pos int, s string) (int, int) {
	i := pos
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	num, _ := strconv.Atoi(s[pos:i])
	return num, i
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	// signs is a stack to signs outside of current parenthese
	signs := []int{1}
	currSign := 1
	i := 0
	sum := 0
	num := 0
	for i < len(s) {
		c := s[i]
		if c >= '0' && c <= '9' {
			num, i = nextNum(i, s)
			// Determine the actual sign by multiplying local sign and outer sign
			sum += currSign * signs[len(signs)-1] * num
		}
		if c == '+' {
			currSign = 1
			i++
		}
		if c == '-' {
			currSign = -1
			i++
		}
		if c == '(' {
			signs = append(signs, currSign*signs[len(signs)-1])
			// Remeber to reset currSign here because a number could follow (
			currSign = 1
			i++
		}
		if c == ')' {
			signs = signs[0 : len(signs)-1]
			// No need to reset currSign here because only operator could follow )
			i++
		}
	}
	return sum
}

func main() {

}
