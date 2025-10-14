package solutions

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	node := result
	rest := 0

	for {
		if l1 == nil {
			l1 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		if l2 == nil {
			l2 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		sum := l1.Val + l2.Val + rest
		rest = 0

		if sum > 9 {
			sum = sum - 10
			rest++
		}

		nextNode := &ListNode{}
		node.Val = sum

		if l1.Next == nil && l2.Next == nil && rest == 0 {
			return result
		}

		node.Next = nextNode
		node = nextNode

		l1 = l1.Next
		l2 = l2.Next
	}
}
