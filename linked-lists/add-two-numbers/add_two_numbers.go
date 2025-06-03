package add_two_numbers

import listnode "github.com/gustavogabana/leetcodegrind/structs"

func AddTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	carry := 0
	dummy := &listnode.ListNode{}
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
		current.Next = &listnode.ListNode{Val: digit}
		current = current.Next
	}
	return dummy.Next
}