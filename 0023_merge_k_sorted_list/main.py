from heapq import heapify, heappush, heappop

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class HeapItem:
    def __init__(self, node):
        self.node = node
        self.val = node.val
    
    def __lt__(self, other):
        return self.val < other.val

class Solution:
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        heap = []
        for node in lists:
            if node is not None:
                heap.append(HeapItem(node))
        heapify(heap)
        dummy = ListNode(0)
        head = dummy
        while len(heap) != 0:
            item = heappop(heap)
            node = item.node
            head.next = node
            head = head.next
            if node.next is not None:
                heappush(heap, HeapItem(node.next))
        return dummy.next