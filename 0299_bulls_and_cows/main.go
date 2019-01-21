package main

import "fmt"

func getHint(secret string, guess string) string {
    counters := map[byte]int{}
    a := 0
    for i := 0; i < len(secret); i++ {
    	if secret[i] == guess[i] {
    		// if postion and char are both correct, add 1 to A
    		a++
    	} else {
    		// else we count how many same char do we have in total that doesn't match the position
    		counters[secret[i]]++
    	}
    }
    b := 0
    for i := 0; i < len(guess); i++ {
    	// if the counter is not zero, and current guess char is different from current secret char
    	// it means we have one char that's not in same position
    	if v, ok := counters[guess[i]]; guess[i] != secret[i] && ok && v > 0 {
    		b += 1
    		counters[guess[i]]--
    	}
    }
    return fmt.Sprintf("%dA%dB", a, b)
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("11", "10"))
	fmt.Println(getHint("1234", "0111"))
	fmt.Println(getHint("1122", "2211"))
}