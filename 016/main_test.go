package main

import (
	"testing"
)

func TestCompress(t *testing.T) {
	//arrange

	testTable := []struct {
		task     string
		expected int
	}{
		{
			task:     "aabbccc", //a2b2c3
			expected: 6,
		},
		{
			task:     "a",
			expected: 1,
		},
		{
			task:     "abbbbbbbbbbbb", //ab12
			expected: 4,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %v, expected: %v", testcase.task, testcase.expected)
		result := compress([]byte(testcase.task))

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect \"%v\", got \"%v\"", testcase.expected, result)

		}
	}

}
