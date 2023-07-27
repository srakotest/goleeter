package main

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	//arrange

	testTable := []struct {
		task     string
		expected string
	}{
		{
			task:     "hello",
			expected: "holle",
		},
		{
			task:     "leetcode",
			expected: "leotcede",
		},
		{
			task:     "programming",
			expected: "prigrammong",
		},
		{
			task:     "leetcode",
			expected: "leotcede",
		},
		{
			task:     "ueo",
			expected: "oeu",
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %s, expected: %s", testcase.task, testcase.expected)
		result := reverseVowels(testcase.task)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %s, got %s", testcase.expected, result)
		}

	}

}
