package main

import "testing"

func TestXIsPalindrome(t *testing.T) {
	//arrange
	testTable := []struct {
		task     string
		expected bool
	}{
		{
			task:     "1",
			expected: true,
		},
		{
			task:     "1-2",
			expected: false,
		},
		{
			task:     "1-2-1",
			expected: true,
		},
		{
			task:     "1-1-2-1",
			expected: false,
		},
		{
			task:     "1-2-3-4-5-6",
			expected: false,
		},
		{
			task:     "1-2-2-1",
			expected: true,
		},
		{
			task:     "1-1",
			expected: true,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Log(testcase.expected, testcase.task)
		result := xisPalindrome(buildList(testcase.task))
		//assert
		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %t, got %t", testcase.expected, result)
		}
	}

}

func TestIsPalindrome(t *testing.T) {
	//arrange
	testTable := []struct {
		task     string
		expected bool
	}{
		{
			task:     "1",
			expected: true,
		},
		{
			task:     "1-2",
			expected: false,
		},
		{
			task:     "1-2-1",
			expected: true,
		},
		{
			task:     "1-1-2-1",
			expected: false,
		},
		{
			task:     "1-2-3-4-5-6",
			expected: false,
		},
		{
			task:     "1-2-2-1",
			expected: true,
		},
		{
			task:     "1-1",
			expected: true,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Log(testcase.expected, testcase.task)
		result := isPalindrome(buildList(testcase.task))
		//assert
		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %t, got %t", testcase.expected, result)
		}
	}

}

func TestIsPalindromeGpt(t *testing.T) {
	//arrange
	testTable := []struct {
		task     string
		expected bool
	}{
		{
			task:     "1",
			expected: true,
		},
		{
			task:     "1-2",
			expected: false,
		},
		{
			task:     "1-2-1",
			expected: true,
		},
		{
			task:     "1-1-2-1",
			expected: false,
		},
		{
			task:     "1-2-3-4-5-6",
			expected: false,
		},
		{
			task:     "1-2-2-1",
			expected: true,
		},
		{
			task:     "1-1",
			expected: true,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Log(testcase.expected, testcase.task)
		result := isPalindrome(buildList(testcase.task))
		//assert
		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %t, got %t", testcase.expected, result)
		}
	}

}
