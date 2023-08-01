package main

import (
	"fmt"
)

/*
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Constraints:

2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

*/

func main() {
	println("Start main")
	task := []int{1, 2, 3, 4}
	fmt.Printf("\"%v\"\n", productExceptSelf(task))
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	curProduct := 1
	for i := len(nums) - 2; i >= 0; i-- {
		curProduct *= nums[i+1]
		result[i] = result[i] * curProduct
	}
	return result
}
