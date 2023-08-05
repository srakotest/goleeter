package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	//arrange

	testTable := []struct {
		task     []int
		expected int
	}{
		{
			task:     []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			task:     []int{1, 1},
			expected: 1,
		},
		{
			task:     []int{0, 100, 0},
			expected: 0,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task, testcase.expected)
		result := maxArea(testcase.task)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
