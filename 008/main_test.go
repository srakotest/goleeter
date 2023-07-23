package main

import (
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
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
			}{"abc", "pqr"},
			expected: "apbqcr",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ab", "pqrs"},
			expected: "apbqrs",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"abcd", "pq"},
			expected: "apbqcd",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"a", "p"},
			expected: "ap",
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %s, %s, expected: %s", testcase.task.word1, testcase.task.word2, testcase.expected)
		result := mergeAlternately(testcase.task.word1, testcase.task.word2)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %s, got %s", testcase.expected, result)
		}
	}

}
func TestNumberOfStepsSlices(t *testing.T) {
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
			}{"abc", "pqr"},
			expected: "apbqcr",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"ab", "pqrs"},
			expected: "apbqrs",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"abcd", "pq"},
			expected: "apbqcd",
		},
		{
			task: struct {
				word1 string
				word2 string
			}{"a", "p"},
			expected: "ap",
		},
	}
	for _, testcase := range testTable {
		//act
		t.Logf("task: %s, %s, expected: %s", testcase.task.word1, testcase.task.word2, testcase.expected)
		result := mergeAlternatelySlices(testcase.task.word1, testcase.task.word2)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %s, got %s", testcase.expected, result)
		}
	}

}
