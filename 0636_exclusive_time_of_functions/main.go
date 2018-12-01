package main

import (
	"fmt"
	"strings"
	"strconv"
)

// due to the weird explanation of end time in the question, we need to add this bias here to make sure endtime is the very end of a second
const bias = 1

func exclusiveTime(n int, logs []string) []int {
    stack := [][]int{
    	[]int{-1, 0}, // add a dummy stack element to make indexing esier later
    }
    res := make([]int, n)
    for _, log := range logs {
    	parts := strings.Split(log, ":")
    	fid, _ := strconv.Atoi(parts[0])
    	op := parts[1]
    	ts, _ := strconv.Atoi(parts[2])
    	top := stack[len(stack) - 1]
    	if op == "start" {
    		// if the log is "start", we need to update the time for last function
    		if top[0] != -1 {
    			res[top[0]] += ts - top[1]
    		}
    		// push this log into stack until we find the end of this function
    		stack = append(stack, []int{fid, ts})
    	} else {
    		// if the log is "end", we need to use the diff between end time and last func startime
    		res[fid] += ts - top[1] + bias
    		stack = stack[:len(stack) - 1]
    		// after poping out last "start", we also update the start time of the one before to reflect the new star time
    		stack[len(stack)-1][1] = ts + bias
    	}
    }
    return res
}

func main() {
	n := 2
	logs := []string{
		"0:start:0",
		"1:start:2",
		"1:end:5",
		"0:end:6",
	}
	fmt.Println(exclusiveTime(n, logs))
}