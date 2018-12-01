package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
    res := []string{}
    cnt := 0
    for i := len(S) - 1; i >= 0; i-- {
    	char := strings.ToUpper(string(S[i]))
    	if char != "-" {
    		if cnt == 0 {
    			res = append(res, "-")
    		}
    		res = append(res, char)
    		cnt = (cnt + 1) % K
    	}
    }
    if len(res) > 0 {
    	res = res[1:]
	    left, right := 0, len(res) - 1
	 	for left < right {
	 		res[left], res[right] = res[right], res[left]
	 		left++
	 		right--
	 	}
	    return strings.Join(res, "")
    }
    return ""
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}