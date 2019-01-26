# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        dummy = ListNode(0)
        head = dummy
        while l1 is not None or l2 is not None:
            if l1 is None:
                head.next = l2
                l2 = l2.next
                break
            elif l2 is None:
                head.next = l1
                l1 = l1.next
                break
            elif l1.val < l2.val:
                head.next = l1
                l1 = l1.next
            else:
                head.next = l2
                l2 = l2.next
            head = head.next
        return dummy.next