class Solution:
    res = []
    mapping = {
        "2": ["a", "b", "c"],
		"3": ["d", "e", "f"],
		"4": ["g", "h", "i"],
		"5": ["j", "k", "l"],
		"6": ["m", "n", "o"],
		"7": ["p", "q", "r", "s"],
		"8": ["t", "u", "v"],
		"9": ["w", "x", "y", "z"],
    }

    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        self.res = []
        if len(digits) != 0:
            self.findCombinations(0, digits, "")
        return self.res
        
    def findCombinations(self, pos, digits, comb):
        if pos >= len(digits):
            self.res.append(comb)
            return
        curr = self.mapping.get(digits[pos])
        if curr is not None:
            for c in curr:
                self.findCombinations(pos+1, digits, comb+c)
