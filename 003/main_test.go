package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	//arrange
	testTable := []struct {
		task     int
		expected []string
	}{
		{
			task:     3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			task:     5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			task:     15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, testcase := range testTable {
		//act
		t.Log(testcase.task, testcase.expected)
		result := fizzBuzz(testcase.task)
		//assert
		for i := 0; i < len(result); i++ {
			if result[i] != testcase.expected[i] {
				t.Errorf("Incorrect Result. Expect %v, got %v", testcase.expected, result)
			}
		}

	}

}
