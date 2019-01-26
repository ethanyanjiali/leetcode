package main

var (
	mapping = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
)

func findCombination(res *[]string, pos int, digits string, comb string) {
	if pos >= len(digits) {
		tmp := make([]byte, len(comb))
		copy(tmp, comb)
		*res = append(*res, string(comb))
		return
	}
	curr, ok := mapping[digits[pos]]
	if ok {
		for _, c := range curr {
			findCombination(res, pos+1, digits, comb+c)
		}
	}
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) != 0 {
		findCombination(&res, 0, digits, "")
	}
	return res
}

func main() {

}
