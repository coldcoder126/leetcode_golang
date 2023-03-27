package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 双指针
	fast := head
	slow := head

	for fast != nil {

		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 快慢指针，如果两者相遇，说明有环，但是相遇的地方不确定
// a = (n-1)(b+c) + c
// 两个指针相遇之后，再走a步，就可以到入环点，即此时再有一个指针p从head处走，必定会在slow处相遇
func detectCycle2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
