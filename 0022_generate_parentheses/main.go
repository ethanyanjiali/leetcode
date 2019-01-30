package main

func findCombinations(left, right int, res *[]string, comb string) {
	if left == 0 && right == 0 && len(comb) != 0 {
		*res = append(*res, comb)
		return
	}
	if left > 0 {
		findCombinations(left-1, right, res, comb+"(")
	}
	if right > 0 && right > left {
		findCombinations(left, right-1, res, comb+")")
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	findCombinations(n, n, &res, "")
	return res
}

func main() {

}
