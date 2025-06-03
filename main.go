package main

import (
	"slices"

	two_sum "github.com/gustavogabana/leetcodegrind/arrays/two-sum"
)

func main() {
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