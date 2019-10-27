package _02

func TwoSum(nums []int, target int) []int {
	var indexes []int
	var sumMap map[int]*int = map[int]*int{}
	for k, num := range nums {
		index := k
		if mapNum := sumMap[num]; mapNum != nil {
			indexes = append(indexes, *sumMap[num])
			indexes = append(indexes, index)
			return indexes
		}
		sumMap[target-num] = &index
	}
	return nil
}
