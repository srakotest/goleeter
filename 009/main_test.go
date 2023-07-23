package main

import (
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	//arrange

	testTable := []struct {
		task struct {
			word1 string
			word2 string
		}
		expected string
	}{
		{
			task: struct {
				word1 string
				word2 string
			}{"ABCABC",
				"ABC"},
			expected: "ABC",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ABABAB", "ABAB"},
			expected: "AB",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"LEET", "CODE"},
			expected: "",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ABCDABC",
				"ABC"},
			expected: "",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ABCDEF",
				"ABC"},
			expected: "",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ABABABAB",
				"ABABABABAB"},
			expected: "AB",
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %s, %s, expected: %s", testcase.task.word1, testcase.task.word2, testcase.expected)
		result := gcdOfStrings(testcase.task.word1, testcase.task.word2)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %s, got %s", testcase.expected, result)
		}
	}

}
