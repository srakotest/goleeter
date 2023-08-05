package main

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	//arrange

	testTable := []struct {
		task     []int
		expected []int
	}{
		{
			task:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			task:     []int{0},
			expected: []int{0},
		},
		{
			task:     []int{1, 1, 4, 3, 12},
			expected: []int{1, 1, 4, 3, 12},
		},
		{
			task:     []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task, testcase.expected)
		result := testcase.task
		moveZeroes(result)

		//assert
		if len(result) != len(testcase.expected) {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)
		}

		for i := 0; i < len(result); i++ {
			if result[i] != testcase.expected[i] {
				t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)
			}
		}
	}

}
