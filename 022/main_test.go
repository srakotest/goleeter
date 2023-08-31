package main

import (
	"testing"
)

func TestMaxVowels(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			str string
			k   int
		}
		expected int
	}{
		{
			task: struct {
				str string
				k   int
			}{
				str: "abciiidef",
				k:   3,
			},
			expected: 3,
		},
		{
			task: struct {
				str string
				k   int
			}{
				str: "aeiou",
				k:   2,
			},
			expected: 2,
		},
		{
			task: struct {
				str string
				k   int
			}{
				str: "leetcode",
				k:   3,
			},
			expected: 2,
		},
		{
			task: struct {
				str string
				k   int
			}{
				str: "r",
				k:   1,
			},
			expected: 0,
		},
		{
			task: struct {
				str string
				k   int
			}{
				str: "a",
				k:   1,
			},
			expected: 1,
		},
		{
			task: struct {
				str string
				k   int
			}{
				str: "qwrqwqwtezxccvbcvbcvxqwrt",
				k:   5,
			},
			expected: 1,
		},
		{
			task: struct {
				str string
				k   int
			}{
				str: "qwrqwqwtezxccvbcvbcvxqwrtee",
				k:   5,
			},
			expected: 2,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task.k, testcase.expected)
		result := maxVowels(testcase.task.str, testcase.task.k)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
