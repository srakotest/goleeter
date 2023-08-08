package main

import (
	"fmt"
)

/*
1679. Max Number of K-Sum Pairs

You are given an integer array nums and an integer k.
In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
Return the maximum number of operations you can perform on the array.

Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= 10^9

Example 1:
Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.

Example 2:
Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.
*/

func main() {
	println("Start main")
	task := []int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}
	fmt.Printf("\"%v\"\n", maxOperations(task, 1))
}

type numsMap map[int]int

func (m numsMap) del(num int) {

	if m[num] > 1 {
		m[num]--
	} else {
		delete(m, num)
	}
}

func maxOperations(nums []int, k int) int {
	n := make(numsMap)
	result := 0
	/*fill the map*/
	for _, v := range nums {
		n[v]++
	}

	/* find pairs */
	for num, cnt := range n {
		isPairExist := true
		if num*2 == k {
			result += cnt / 2
			n[num] = cnt % 2
			continue
		}
		for cnt > 0 && isPairExist {

			isPairExist = false
			if num <= k {
				_, isPairExist = n[k-num]
				if isPairExist {
					result++
					n.del(num)
					cnt--
					n.del(k - num)
					if k == num {
						cnt--
					}
				}
			}
		}
	}

	return result
}
