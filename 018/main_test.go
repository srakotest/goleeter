package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			word1 string
			word2 string
		}
		expected bool
	}{
		{
			task: struct {
				word1 string
				word2 string
			}{"abc",
				"ahbgdc"},
			expected: true,
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"axc",
				"ahbgdc"},
			expected: false,
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ace", "arcade"},
			expected: true,
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"", ""},
			expected: true,
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"", "qweqweqw"},
			expected: true,
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"sadasdsad", ""},
			expected: false,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %s, %s, expected: %t", testcase.task.word1, testcase.task.word2, testcase.expected)
		result := isSubsequence(testcase.task.word1, testcase.task.word2)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %t, got %t", testcase.expected, result)
		}
	}

}
