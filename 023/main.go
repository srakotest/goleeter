package main

import (
	"fmt"
)

/*
1004. Max Consecutive Ones III

Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

Constraints:
1 <= nums.length <= 10^5
nums[i] is either 0 or 1.
0 <= k <= nums.length

Example 1:
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

[0,0,1,1,1,0,0]
k=0
result=3

*/

func main() {
	println("Start main")
	task := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Printf("\"%v\"\n", longestOnes(task, 2))
}

func longestOnes(nums []int, k int) int {

	//vars
	maxWindow := 0
	currentWindow := 0
	positionFront := 0

	for positionFront < len(nums) {
		//move front
		for positionFront < len(nums) && (nums[positionFront] == 1 || k > 0) {
			if nums[positionFront] == 0 {
				k--
			}
			positionFront++
			currentWindow++
		}

		// save max
		if currentWindow > maxWindow {
			maxWindow = currentWindow
		}

		// move back
		if k == 0 {
			for k == 0 && currentWindow > 0 {
				if nums[positionFront-currentWindow] == 0 {
					k++
				}
				currentWindow--
			}
		}

		if k == 0 && positionFront < len(nums) {
			positionFront++
		}

	}

	return maxWindow
}
