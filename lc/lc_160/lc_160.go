package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表相交，双指针遍历
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != nil || pb != nil {
		if pa == pb {
			return pa
		}
		if pa == nil && pb != nil {
			pa = headB
		}
		if pa != nil && pb == nil {
			pb = headA
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}
