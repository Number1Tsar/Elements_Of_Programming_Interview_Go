package linkedList

/*
	Leetcode Question 92
	Reverse Linked List II
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	index := 0
	start := dummyHead
	for start != nil {
		if left == index+1 {
			break
		}
		start = start.Next
		index += 1
	}
	var p, q, r *ListNode
	p = nil
	q = start.Next
	index += 1
	for q != nil && index <= right {
		r = q.Next
		q.Next = p
		p = q
		q = r
		index += 1
	}
	start.Next.Next = q
	start.Next = p
	return dummyHead.Next
}
