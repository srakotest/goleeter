package main

import "testing"

func TestMiddleNode(t *testing.T) {
	//arrange
	testTable := []struct {
		task     string
		expected string
	}{
		{
			task:     "1",
			expected: "1",
		},
		{
			task:     "1-2",
			expected: "2",
		},
		{
			task:     "1-2-1",
			expected: "2-1",
		},
		{
			task:     "1-1-2-1",
			expected: "2-1",
		},
		{
			task:     "1-2-3-4-5-6",
			expected: "4-5-6",
		},
		{
			task:     "1-2-2-1",
			expected: "2-1",
		},
		{
			task:     "1-2-3-4-5",
			expected: "3-4-5",
		},
	}

	for _, testcase := range testTable {
		//act
		t.Log(testcase.task, testcase.expected)
		result := printList(middleNode(buildList(testcase.task)))
		//assert
		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %s, got %s", testcase.expected, result)
		}
	}

}
