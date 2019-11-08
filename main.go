package main

import (
	"fmt"
	_03 "leetcode/lesson/003"
)

func main() {
	//001
	//ins := []int{1,1,1,2,3,4,5,5,6,6,6,7,7,7,7,7,7}
	//fmt.Println(_01.RemoveDuplicatesFromSortedArray(ins))
	//002
	//fmt.Println(_02.TwoSum([]int{2, 11, 7, 15}, 9))
	/**
	  示例：
	  输入：(9 -> 4 -> 3) + (5 -> 6 -> 4)
	  输出：4 -> 1 -> 8
	  原因：349 + 465 = 814
	*/
	// 003
	l1 := &_03.ListNode{Val: 9}
	l1.Next = &_03.ListNode{Val: 4}
	l1.Next.Next = &_03.ListNode{Val: 3}
	//349 + 465 = 814

	l2 := &_03.ListNode{Val: 5}
	l2.Next = &_03.ListNode{Val: 6}
	l2.Next.Next = &_03.ListNode{Val: 4}

	l3 := _03.AddTwoNumbers(l1, l2)
	for {
		fmt.Println(l3.Val)
		if l3.Next == nil {
			break
		}
		l3 = l3.Next
	}
}
