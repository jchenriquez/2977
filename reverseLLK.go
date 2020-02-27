package reversellk

// ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode, k int) (nHead *ListNode, tail *ListNode) {
	var prev *ListNode
	var nextT *ListNode
	next := head

	for i := 0; i < k; i++ {
		nextT = next.Next
		next.Next = prev

		if prev == nil {
			tail = next
		}

		prev = next
		next = nextT
	}

	nHead = prev

	return
}

// ReverseKGroup will reverse Linked List by k batches.
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	voidNode := &ListNode{-1, nil}
	currHead := voidNode
	curr := voidNode
	voidNode.Next = head
	var count int

	for curr != nil {
		curr = curr.Next
		count++

		if count == k && curr == nil {
			next := curr.Next

			head, tail := reverse(currHead.Next, k)

			currHead.Next = head
			tail.Next = next
			currHead = tail
			curr = currHead
			count = 0
		}
	}

	return voidNode.Next
}
