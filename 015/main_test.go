package main

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	//arrange

	testTable := []struct {
		task     []int
		expected bool
	}{
		{
			task:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			task:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			task:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
		{
			task:     []int{20, 100, 10, 12, 5, 13},
			expected: true,
		},
		{
			task:     []int{50, 51, 9, 10, 8, 9, 7, 8, 6, 12, 0, 1},
			expected: true,
		},
		{
			task:     []int{1},
			expected: false,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task, testcase.expected)
		result := increasingTriplet(testcase.task)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
