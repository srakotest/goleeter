package main

import (
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	//arrange
	testTable := []struct {
		mat      [][]int
		k        int
		expected []int
	}{
		{
			mat: [][]int{
				{1, 1},
				{1, 0},
			},
			k:        1,
			expected: []int{1},
		},
		{
			mat: [][]int{
				[]int{1, 1, 0},
				[]int{1, 0, 0},
				[]int{1, 1, 1},
				[]int{0, 0, 0},
			},
			k:        3,
			expected: []int{3, 1, 0},
		},
		{
			mat: [][]int{
				[]int{1, 0, 0, 0},
				[]int{1, 1, 1, 1},
				[]int{1, 0, 0, 0},
				[]int{1, 0, 0, 0},
			},
			k:        2,
			expected: []int{0, 2},
		},
		{
			mat: [][]int{
				[]int{1, 0, 0, 0, 0},
				[]int{1, 1, 1, 1, 1},
				[]int{0, 0, 0, 0, 0},
				[]int{1, 0, 0, 0, 1},
			},
			k:        4,
			expected: []int{2, 0, 3, 1},
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("mat: %v, k: %d, expected: %v", testcase.mat, testcase.k, testcase.expected)
		result := kWeakestRows(testcase.mat, testcase.k)

		//assert
		if len(result) != len(testcase.expected) {
			t.Errorf("Incorrect Result length. Expect %v, got %v", testcase.expected, result)
		}

		for i := 0; i < len(result) && i < len(testcase.expected); i++ {
			if result[i] != testcase.expected[i] {
				t.Errorf("Incorrect Result. Expect %v, got %v", testcase.expected, result)
			}
		}
	}

}
