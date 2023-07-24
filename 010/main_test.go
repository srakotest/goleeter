package main

import (
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			candies      []int
			extraCandies int
		}
		expected []bool
	}{
		{
			task: struct {
				candies      []int
				extraCandies int
			}{[]int{2, 3, 5, 1, 3},
				3},
			expected: []bool{true, true, true, false, true},
		},
		{
			task: struct {
				candies      []int
				extraCandies int
			}{[]int{4, 2, 1, 1, 2},
				1},
			expected: []bool{true, false, false, false, false},
		},
		{
			task: struct {
				candies      []int
				extraCandies int
			}{[]int{12, 1, 12},
				10},
			expected: []bool{true, false, true},
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, %d, expected: %v", testcase.task.candies, testcase.task.extraCandies, testcase.expected)
		result := kidsWithCandies(testcase.task.candies, testcase.task.extraCandies)
		//assert
		for i := 0; i < len(result); i++ {
			if result[i] != testcase.expected[i] {
				t.Errorf("Incorrect Result. Expect %v, got %v", testcase.expected, result)
			}
		}
	}

}
