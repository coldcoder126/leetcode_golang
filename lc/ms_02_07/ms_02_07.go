package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// A遍历完再遍历B，B遍历完再遍历A。循环
	pa := headA
	pb := headB
	for pa != nil || pb != nil {
		pa = pa.Next
		pb = pb.Next

		if pa == nil && pb != nil {
			pa = headB
		}
		if pb == nil && pa != nil {
			pb = headA
		}
	}
	return nil
}
