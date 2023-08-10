package main

import (
	"fmt"
)

/*
643. Maximum Average Subarray I

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than 10^-5 will be accepted.

Constraints:

n == nums.length
1 <= k <= n <= 10^5
-10^4 <= nums[i] <= 10^4


Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2
Input: nums = [5], k = 1
Output: 5.00000
*/

func main() {
	println("Start main")
	task := []int{1, 12, -5, -6, 50, 3}
	fmt.Printf("\"%v\"\n", findMaxAverage(task, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	var maxAverage, curAverage float64
	maxAverage = float64(0 - k*10000)
	for i, v := range nums {
		if i >= k {
			curAverage -= float64(nums[i-k]) / float64(k)
		}
		curAverage += float64(v) / float64(k)
		if curAverage > maxAverage && i >= k-1 {
			maxAverage = curAverage
		}
	}

	return maxAverage
}
