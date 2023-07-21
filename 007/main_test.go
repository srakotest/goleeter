package main

import (
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	//arrange

	testTable := []struct {
		task     [][]int
		expected int
	}{
		{
			task:     [][]int{{1, 2, 3}, {3, 2, 1}},
			expected: 6,
		},
		{
			task:     [][]int{{1, 5}, {7, 3}, {3, 5}},
			expected: 10,
		},
		{
			task:     [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			expected: 17,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %d", testcase.task, testcase.expected)
		result := maximumWealth(testcase.task)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
		}
	}
}
