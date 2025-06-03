package main

import (
	two_sum "github.com/gustavogabana/leetcodegrind/arrays/two-sum"
	add_two_numbers "github.com/gustavogabana/leetcodegrind/linked-lists/add-two-numbers"
	listnode "github.com/gustavogabana/leetcodegrind/structs"
)

func call_two_sum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := two_sum.TwoSum(nums, target)
	println("Resultado esperado: [0,1], ", result)
}

func call_add_two_numbers() {
	l1 := &listnode.ListNode{
		Val: 2,
		Next: &listnode.ListNode{
			Val: 4,
			Next: &listnode.ListNode{
				Val: 3,
			},
		},
	}
	l2 := &listnode.ListNode{
		Val: 5,
		Next: &listnode.ListNode{
			Val: 6,
			Next: &listnode.ListNode{
				Val: 4,
			},
		},
	}
	result := add_two_numbers.AddTwoNumbers(l1, l2)
	println("Resultado esperado: 7 -> 0 -> 8, ", result)
}

func main() {
	call_two_sum()
	call_add_two_numbers()
}
