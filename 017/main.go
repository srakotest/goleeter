package main

import (
	"fmt"
)

/*
283. Move Zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Constraints:

1 <= nums.length <= 10^4
-2^31 <= nums[i] <= 2^31 - 1

Follow up: Could you minimize the total number of operations done?


Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]

*/

func main() {
	println("Start main")
	task := []int{0, 1, 0, 3, 12}
	moveZeroes(task)
	fmt.Printf("\"%v\"\n", task)
}

func moveZeroes(nums []int) {
	zerosCounter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zerosCounter++
		} else {
			nums[i-zerosCounter] = nums[i]
		}
	}
	for i := len(nums) - zerosCounter; i < len(nums); i++ {
		nums[i] = 0
	}
}

/* first version
func moveZeroes(nums []int) {
	endNonZero := len(nums) - 1
	i := 0
	for i <= endNonZero {
		if nums[i] == 0 {
			for m := i; m < endNonZero; m++ {
				nums[m], nums[m+1] = nums[m+1], nums[m]
			}
			endNonZero--
		} else {
			i++
		}
	}
}*/
