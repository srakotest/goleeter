package main

import (
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	//arrange

	testTable := []struct {
		task     []int
		expected int
	}{
		{
			task:     []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			task:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			task:     []int{1, 1, 1},
			expected: 2,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task, testcase.expected)
		result := longestSubarray(testcase.task)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
