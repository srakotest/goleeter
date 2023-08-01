package main

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	//arrange

	testTable := []struct {
		task     []int
		expected []int
	}{
		{
			task:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			task:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			task:     []int{5, 10},
			expected: []int{10, 5},
		},
		{
			task:     []int{-1, 1, 0, -3, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task, testcase.expected)
		result := productExceptSelf(testcase.task)

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
