package main

import (
	"testing"
)

func TestLongestOnes(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			nums []int
			k    int
		}
		expected int
	}{
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			expected: 6,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3,
			},
			expected: 10,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{0},
				k:    0,
			},
			expected: 0,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{1},
				k:    0,
			},
			expected: 1,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{0, 0, 0, 0, 0},
				k:    5,
			},
			expected: 5,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{0, 0, 1, 1, 1, 0, 0},
				k:    0,
			},
			expected: 3,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task.k, testcase.expected)
		result := longestOnes(testcase.task.nums, testcase.task.k)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
