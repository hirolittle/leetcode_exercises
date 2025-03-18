package twoSum

// 两数之和
// https://leetcode.cn/problems/two-sum/description/

func twoSum(nums []int, target int) []int {

	// 方法一: 暴力枚举
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return nil

	// 方法二：哈希表
	hashTable := make(map[int]int)
	for i, num := range nums {
		if j, ok := hashTable[target-num]; ok {
			return []int{j, i}
		}
		hashTable[num] = i
	}
	return nil
}
