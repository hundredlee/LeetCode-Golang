package _01

import "fmt"

func RemoveDuplicatesFromSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	/**
	j 代表每个不重复数字的位置，当前元素与前一个元素相同时，j就不需要移动。
	*/
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums[:j])
	return j
}
