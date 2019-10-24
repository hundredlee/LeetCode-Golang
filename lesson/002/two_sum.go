package _02

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func TwoSum(nums []int, target int) []int {
	var indexes []int
	var sumMap map[int]*int = map[int]*int{}
	for k, num := range nums {
		index := k
		if mapNum := sumMap[num]; mapNum != nil {
			indexes = append(indexes,*sumMap[num])
			indexes = append(indexes,index)
			return indexes
		}
		sumMap[target-num] = &index
	}
	return nil
}
