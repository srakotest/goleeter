package main

import (
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			flowerbed    []int
			extraFlowers int
		}
		expected bool
	}{
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{1, 0, 0, 0, 1},
				1},
			expected: true,
		},
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{1, 0, 0, 0, 1},
				2},
			expected: false,
		},
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{1, 0},
				0},
			expected: true,
		},
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{0},
				1},
			expected: true,
		},
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{0, 0},
				1},
			expected: true,
		},
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{0, 1},
				1},
			expected: false,
		},
		{
			task: struct {
				flowerbed    []int
				extraFlowers int
			}{[]int{0, 1, 0, 1, 0, 1, 0, 0},
				1},
			expected: true,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, %d, expected: %v", testcase.task.flowerbed, testcase.task.extraFlowers, testcase.expected)
		result := canPlaceFlowers(testcase.task.flowerbed, testcase.task.extraFlowers)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %v, got %v", testcase.expected, result)
		}

	}

}
