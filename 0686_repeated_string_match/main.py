class Solution:
    def repeatedStringMatch(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: int
        """
        times = math.ceil(len(B) / len(A))
        match = A * times
        if B in match:
            return times
        match += A
        times += 1
        if B in match:
            return times
        return -1