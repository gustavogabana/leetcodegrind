package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	current := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		value1 := 0
		if l1 != nil {
			value1 = l1.Val
			l1 = l1.Next
		}
		value2 := 0
		if l2 != nil {
			value2 = l2.Val
			l2 = l2.Next
		}
		sum := value1 + value2 + carry
		digit := sum % 10
		carry = sum / 10
		current.Next = &ListNode{Val: digit}
		current = current.Next
	}
	return dummy.Next
}