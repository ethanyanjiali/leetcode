package main

import (
	"fmt"
	"strconv"
	"strings"
)

func say(s string) string {
	last := []rune(s)[0]
	var res strings.Builder
	count := 0x30 // utf-8 '0'
	for _, r := range s {
		if r != last {
			res.WriteRune(rune(count))
			res.WriteRune(last)
			last = r
			count = 0x31 // utf-8 '1'
		} else {
			count = count + 1
		}
	}
	res.WriteRune(rune(count))
	res.WriteRune(last)
	return res.String()
}

func countAndSay(n int) string {
	res := strconv.Itoa(1)
	// when n=1, we just output 1
	for i := 2; i <= n; i = i + 1 {
		res = say(res)
	}
	return res
}

func main() {
	fmt.Println(countAndSay(3))
}
