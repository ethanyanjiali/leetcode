from collections import deque


class Solution:
    def findCheapestPrice(self, n, flights, src, dst, K):
        """
        :type n: int
        :type flights: List[List[int]]
        :type src: int
        :type dst: int
        :type K: int
        :rtype: int
        """
        graph = self.buildGraph(flights)
        q = deque([(src, 0)])
        # stops start from -1 because K could be 0 (non-stop)
        stops = -1
        visited = {}
        cheapest = -1
        # we need to go to the dst to verify result, so <= K
        while len(q) > 0 and stops <= K:
            size = len(q)
            for i in range(size):
                start, total = q.popleft()
                if start == dst and (cheapest == -1 or cheapest > total):
                    cheapest = total
                    continue
                toMap = graph[start]
                for to, cost in toMap.items():
                    if to in visited and visited[to] > total + cost:
                        continue
                    q.append((to, total+cost))
                    visited[to] = total+cost
            stop += 1
        return cheapest

    def buildGraph(self, flights):
        graph = {}
        for flight in flights:
            src, dst, cost = flight
            if src in graph:
                graph[src] = {}
            cmap = graph[src]
            cmap[dst] = cost
            graph[src] = cmap
        return graph
                
            