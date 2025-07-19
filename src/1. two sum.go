// https://leetcode.com/problems/two-sum/

package main

func twoSum(nums []int, target int) []int {
	complement_index := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, exists := complement_index[num]; exists {
			return []int{i, index}
		}

		complement_index[complement] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	println(result[0], result[1])
}
