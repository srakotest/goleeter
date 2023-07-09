package main

import (
	"math"
	"testing"
)

type tCanConstructInputs struct {
	ransomNote string
	magazine   string
}
type tTestCase struct {
	task     tCanConstructInputs
	expected bool
}

func TestCanConstruct(t *testing.T) {
	//arrange
	testTable := []tTestCase{
		{
			task:     tCanConstructInputs{ransomNote: "aaa", magazine: "aaaa"},
			expected: true,
		},
		{
			task:     tCanConstructInputs{"aaa", "aa"},
			expected: false,
		},
		{
			task:     tCanConstructInputs{"aaab", "aaaaqwe"},
			expected: false,
		},
		{
			task:     tCanConstructInputs{"p", "qwertyuiop"},
			expected: true,
		},
		{
			task:     tCanConstructInputs{"s", "s"},
			expected: true,
		},
		{
			task:     tCanConstructInputs{"a", "z"},
			expected: false,
		},
	}

	{
		var x string
		for i := 0; i < int(math.Pow10(5)); i++ {
			x += "a"
		}

		testTable = append(testTable, tTestCase{tCanConstructInputs{x, x}, true})
	}

	for _, testcase := range testTable {
		//act
		t.Log(testcase.expected, testcase.task)
		result := canConstruct(splitStruct(testcase.task))
		//assert
		if result != testcase.expected {
			t.Errorf("Incorrect Result. Expect %t, got %t", testcase.expected, result)
		}
	}

}

func splitStruct(s tCanConstructInputs) (string, string) {
	return s.ransomNote, s.magazine
}

/*
func addCase(sl []tTestCase, newCase tTestCase) {
	sl = append(sl, newCase)
}
*/
