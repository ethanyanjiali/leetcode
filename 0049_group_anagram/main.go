package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	group := map[string][]string{}
	res := [][]string{}
	for _, s := range strs {
		str := []byte(s)
		sort.Slice(str, func(i, j int) bool {
			return str[i] <= str[j]
		})
		key := string(str)
		_, ok := group[key]
		if ok {
			group[key] = append(group[key], s)
		} else {
			group[key] = []string{s}
		}
	}
	for _, v := range group {
		res = append(res, v)
	}
	return res
}

func main() {

}
