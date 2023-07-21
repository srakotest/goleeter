package main

import (
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	//arrange

	testTable := []struct {
		task     int
		expected int
	}{
		{
			task:     14,
			expected: 6,
		},
		{
			task:     8,
			expected: 4,
		},
		{
			task:     123456,
			expected: 22,
		},
		{
			task:     123,
			expected: 12,
		},
		{
			task:     1000000,
			expected: 26,
		},
		{
			task:     987654,
			expected: 27,
		},
		{
			task:     994991,
			expected: 33,
		},
		{
			task:     524287,
			expected: 37,
		},
	}

	for _, testcase := range testTable {
		//act
		t.Logf("task: %d, expected: %d", testcase.task, testcase.expected)
		result := numberOfSteps(testcase.task) //numberOfSteps(testcase.task)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
		}
	}

	/*bits count*/
	t.Logf("----------------  BITS  -------------")
	for _, testcase := range testTable {
		//act
		t.Logf("task: %d, expected: %d", testcase.task, testcase.expected)
		result := bitnumberOfSteps(testcase.task) //numberOfSteps(testcase.task)

		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
		}
	}

	/* bit shift */
	t.Logf("----------------  BITSHIFT  -------------")
	for _, testcase := range testTable {
		//act
		t.Logf("task: %d, expected: %d", testcase.task, testcase.expected)
		result := numberOfStepsByBits(testcase.task)
		//assert

		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
		}
	}
}

func TestNumberOfStepsLoad(t *testing.T) {
	//arrange

	testTable := []struct {
		task     int
		expected int
	}{
		{
			task:     14,
			expected: 6,
		},
		{
			task:     8,
			expected: 4,
		},
		{
			task:     123456,
			expected: 22,
		},
		{
			task:     123,
			expected: 12,
		},
		{
			task:     1000000,
			expected: 26,
		},
		{
			task:     987654,
			expected: 27,
		},
		{
			task:     994991,
			expected: 33,
		},
		{
			task:     524287,
			expected: 37,
		},
	}

	for i := 0; i < 100000000; i++ {
		for _, testcase := range testTable {
			//act
			result := numberOfSteps(testcase.task) //numberOfSteps(testcase.task)

			//assert

			if result != testcase.expected {
				t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
			}
		}
	}
}

func TestNumberOfStepsLoadBits(t *testing.T) {
	//arrange

	testTable := []struct {
		task     int
		expected int
	}{
		{
			task:     14,
			expected: 6,
		},
		{
			task:     8,
			expected: 4,
		},
		{
			task:     123456,
			expected: 22,
		},
		{
			task:     123,
			expected: 12,
		},
		{
			task:     1000000,
			expected: 26,
		},
		{
			task:     987654,
			expected: 27,
		},
		{
			task:     994991,
			expected: 33,
		},
		{
			task:     524287,
			expected: 37,
		},
	}

	for i := 0; i < 100000000; i++ {
		for _, testcase := range testTable {
			//act
			result := bitnumberOfSteps(testcase.task) //numberOfSteps(testcase.task)

			//assert

			if result != testcase.expected {
				t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
			}
		}
	}
}

func TestNumberOfStepsLoadByBits(t *testing.T) {
	//arrange

	testTable := []struct {
		task     int
		expected int
	}{
		{
			task:     14,
			expected: 6,
		},
		{
			task:     8,
			expected: 4,
		},
		{
			task:     123456,
			expected: 22,
		},
		{
			task:     123,
			expected: 12,
		},
		{
			task:     1000000,
			expected: 26,
		},
		{
			task:     987654,
			expected: 27,
		},
		{
			task:     994991,
			expected: 33,
		},
		{
			task:     524287,
			expected: 37,
		},
	}

	for i := 0; i < 100000000; i++ {
		for _, testcase := range testTable {
			//act
			result := numberOfStepsByBits(testcase.task) //numberOfSteps(testcase.task)

			//assert

			if result != testcase.expected {
				t.Errorf("Incorrect Result. Expect %d, got %d", testcase.expected, result)
			}
		}
	}
}
