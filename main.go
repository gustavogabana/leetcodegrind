package main

import (
	"slices"

	two_sum "github.com/gustavogabana/leetcodegrind/arrays/two-sum"
)

func call_two_sum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result_expected := []int{0, 1}
	result := two_sum.TwoSum(nums, target)
	if slices.Equal(result_expected, result) {
		println("TWO SUM PASSED")
	} else {
		println("TWO SUM NOT PASSED")
	}
}

func call_add_two_numbers() {
	
}

func main() {
	call_two_sum()
}