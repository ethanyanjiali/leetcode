package main

import (
	"fmt"
)

func getKey(key1, key2 byte) string {
	keyA := string(key1)
	keyB := string(key2)
	return keyA + keyB
}

func findNextLevels(bottom string, options map[string]map[byte]bool) []string {
	level := []string{""}
	for i := 0; i < len(bottom) - 1; i++ {
		char1 := bottom[i]
		char2 := bottom[i + 1]
		option := options[getKey(char1, char2)]
		b := []string{}
		for _, a := range level {
			for key := range option {
				b = append(b, a + string(key))
			}
		}
		level = b
	}
	return level
}

func buildPyramidByLevel(bottom string, options map[string]map[byte]bool) bool {
	if len(bottom) == 1 {
		return true
	}
	nextBottoms := findNextLevels(bottom, options)
	for _, nextBottom := range nextBottoms {
		if buildPyramidByLevel(nextBottom, options) == true {
			return true
		}
	}
	return false
}

func pyramidTransition1(bottom string, allowed []string) bool {
	options := map[string]map[byte]bool{}
	for _, tuple := range allowed {
		key := getKey(tuple[0], tuple[1])
		if _, ok := options[key]; !ok {
			options[key] = map[byte]bool{}
		}
		options[key][tuple[2]] = true
	}
	return buildPyramidByLevel(bottom, options)
}

func buildPyramidByBlock(bottom string, top string, options map[string]map[string]bool) bool {
	if len(bottom) == 2 && len(top) == 1 {
		return true
	}
	if len(bottom) - 1 == len(top) {
		return buildPyramidByBlock(top, "", options)
	}
	option := options[getKey(bottom[len(top)], bottom[len(top)+1])]
	for block := range option {
		if buildPyramidByBlock(bottom, top+block, options) {
			return true
		}
	}
	return false
}

func pyramidTransition2(bottom string, allowed []string) bool {
	options := map[string]map[string]bool{}
	for _, tuple := range allowed {
		key := getKey(tuple[0], tuple[1])
		if _, ok := options[key]; !ok {
			options[key] = map[string]bool{}
		}
		options[key][string(tuple[2])] = true
	}
	return buildPyramidByBlock(bottom, "", options)
}

func main() {
	fmt.Println(pyramidTransition2(
		"XYZ",
		[]string{"XYD", "YZE", "DEA", "FFF"},
	))
	fmt.Println(pyramidTransition2(
		"XXYX",
		[]string{"XXX", "XXY", "XYX", "XYY", "YXZ"},
	))
	fmt.Println(pyramidTransition2(
		"AAAA",
		[]string{"AAA",},
	))
}


