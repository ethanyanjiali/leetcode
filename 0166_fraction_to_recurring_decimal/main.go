package main

import (
	"strconv"
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	sign := ""
	if numerator < 0 && denominator > 0 {
		sign = "-"
		numerator = -numerator
	}
	if denominator < 0 && numerator > 0 {
		sign = "-"
		denominator = -denominator
	}
    integer := strconv.Itoa(numerator / denominator)
    fractional := ""
    remainder := numerator % denominator
    appeared := map[int]int{}
    for remainder != 0 {
    	if index, ok := appeared[remainder]; ok {
    		repeat := fractional[index:]
    		fractional = fractional[:index]+"("+repeat+")"
    		break
    	} else {
    		appeared[remainder] = len(fractional)
    		digit := remainder * 10 / denominator
    		remainder = remainder * 10 % denominator
    		fractional += strconv.Itoa(digit)
    	}
    }
    if len(fractional) > 0 {
    	return sign + integer + "." + fractional
    }
    return sign + integer
}

func main() {
	fmt.Println(fractionToDecimal(1, 7))
}