// https://leetcode.com/problems/two-sum/

package src

import "testing"

func twoSum(nums []int, target int) []int {
	complement_index := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, exists := complement_index[num]; exists {
			return []int{index, i}
		}

		complement_index[complement] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)

	if result == nil {
		t.Error("Expected result, got nil")
		return
	}

	t.Logf("Result: [%d, %d]", result[0], result[1])

	// Verify the result is correct
	if nums[result[0]]+nums[result[1]] != target {
		t.Errorf("Expected sum to be %d, got %d", target, nums[result[0]]+nums[result[1]])
	}
}
