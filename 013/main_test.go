package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	//arrange

	testTable := []struct {
		task     string
		expected string
	}{
		{
			task:     "the sky is blue",
			expected: "blue is sky the",
		},
		{
			task:     "  hello world  ",
			expected: "world hello",
		},
		{
			task:     "a good   example",
			expected: "example good a",
		},
		{
			task:     "pneumonoultramicroscopicsilicovolcanoconiosis",
			expected: "pneumonoultramicroscopicsilicovolcanoconiosis",
		},
		{
			task:     "wonderful ",
			expected: "wonderful",
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: \"%s\", expected: \"%s\"", testcase.task, testcase.expected)
		result := reverseWords(testcase.task)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%s\", got \"%s\"", testcase.expected, result)
		}

	}

}
