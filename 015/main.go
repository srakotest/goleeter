package main

import (
	"fmt"
)

/*
334. Increasing Triplet Subsequence

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that
 i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

Constraints:
1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1

Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?


Example 1:
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.

Example 3:
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.

*/

func main() {
	println("Start main")
	task := []int{1, 2, 3, 4, 5}
	fmt.Printf("\"%v\"\n", increasingTriplet(task))
}

func increasingTriplet(nums []int) bool {

	if len(nums) < 3 {
		return false
	}

	left := nums[0]
	mid := nums[1]
	stage := 'l'

	for i := 1; i < len(nums); i++ {
		switch stage {
		case 'l':
			if nums[i] < left {
				left = nums[i]
			} else if nums[i] > left {
				stage = 'm'
				mid = nums[i]
			}
		case 'm':
			if nums[i] > mid {
				return true
			} else if nums[i-1] < left && nums[i] > nums[i-1] && i < len(nums)-1 {
				left = nums[i-1]
				mid = nums[i]
			} else if nums[i] < mid && nums[i] > left {
				mid = nums[i]
			}
		}
	}

	return false
}
