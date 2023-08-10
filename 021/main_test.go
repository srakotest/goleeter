package main

import (
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			nums []int
			k    int
		}
		expected float64
	}{
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			expected: 12.75,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{5},
				k:    1,
			},
			expected: 5.0,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{5},
				k:    1,
			},
			expected: 5.0,
		},
		{
			task: struct {
				nums []int
				k    int
			}{
				nums: []int{-1},
				k:    1,
			},
			expected: -1.0,
		},
	}

	for _, testcase := range testTable {
		//act
		//t.Logf("task: %v, expected: %v", testcase.task.k, testcase.expected)
		result := findMaxAverage(testcase.task.nums, testcase.task.k)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
