package _03

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: 0}
	cursor := root
	var carry int = 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		var l1Val int = 0
		var l2Val int = 0
		var sumVal int = 0
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		sumVal = l1Val + l2Val + carry
		carry = sumVal / 10

		sumNode := &ListNode{Val: sumVal % 10}
		cursor.Next = sumNode
		cursor = sumNode

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}
