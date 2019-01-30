class Solution:
    res = []

    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        self.res = []
        self.findCombinations(n, n, "")
        return self.res
    
    def findCombinations(self, left, right, comb):
        if left == 0 and right == 0 and len(comb) != 0 {
            self.res.append(comb)
            return
        }
        if left > 0:
            self.findCombinations(left-1, right, comb+"(")
        if right > 0 and right > left:
            self.findCombinations(left, right-1, comb+")")