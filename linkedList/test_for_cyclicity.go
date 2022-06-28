package linkedList

/*
	LeetCode question no 141
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return false
}

/*
	LeetCode question no 142. (Finding the beginning of the cycle, if it exists
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}
	return nil
}
