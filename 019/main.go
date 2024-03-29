package main

import (
	"fmt"
)

/*
11. Container With Most Water

You are given an integer array height of length n. There are n vertical lines drawn such that
the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Constraints:
n == height.length
2 <= n <= 10^5
0 <= height[i] <= 10^4

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1



*/

func main() {
	println("Start main")
	task := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("\"%v\"\n", maxArea(task))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxS := 0
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}

		s := h * (r - l)

		if s > maxS {
			maxS = s
		}
		if height[l] > h {
			r--
		} else {
			l++
		}
	}
	return maxS
}

/*
first version

func maxArea(height []int) int {
	maxS := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			s := (j - i) * h
			if s > maxS {
				maxS = s
			}
		}
	}
	return maxS
}

*/
