package main

import (
	"fmt"
)

/*
1493. Longest Subarray of 1's After Deleting One Element

Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array.
Return 0 if there is no such subarray.

Constraints:
1 <= nums.length <= 10^5
nums[i] is either 0 or 1.

Example 1:
Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.

Example 2:
Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].

Example 3:
Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.

*/

func main() {
	println("Start main")
	task := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Printf("\"%v\"\n", longestSubarray(task))
}

func longestSubarray(nums []int) int {
	rightPointer, leftPointer, maxWindow, zeros := 0, 0, 0, 1

	for rightPointer < len(nums) {
		if nums[rightPointer] == 0 {
			zeros--
		}

		for leftPointer <= rightPointer && zeros < 0 {
			if nums[leftPointer] == 0 {
				zeros++
			}
			leftPointer++
		}

		if rightPointer-leftPointer > maxWindow {
			maxWindow = rightPointer - leftPointer
		}

		rightPointer++
	}

	return maxWindow
}
