package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	var (
		prev    *ListNode
		current *ListNode
	)

	current = head
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	return prev
}

func reverse(prev, current *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	tmp := current.Next
	current.Next = prev
	return reverse(current, tmp)
}

func reverseList2(head *ListNode) *ListNode {
	return reverse(nil, head)
}
