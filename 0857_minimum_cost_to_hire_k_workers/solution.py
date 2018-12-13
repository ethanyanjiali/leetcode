import heapq

class Solution:
    def mincostToHireWorkers(self, quality, wage, K):
        """
        :type quality: List[int]
        :type wage: List[int]
        :type K: int
        :rtype: float
        """
        ratios = [(w / q, q) for w, q in zip(wage, quality)]
        ratios.sort(key=lambda x: x[0])
        heap = []
        qsum = 0
        min_total = float('inf')
        for ratio in ratios:
        	heapq.heappush(heap, -ratio[1])
        	qsum += ratio[1]
        	print(qsum, ratio)
        	if len(heap) == K + 1:
        		maxq = heapq.heappop(heap)
        		qsum += maxq
        	print(qsum, ratio)
        	if len(heap) == K:
        		min_total = min(min_total, qsum * ratio[0])
        return min_total


