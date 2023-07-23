package main

import (
	"fmt"
)

/*
1071. Greatest Common Divisor of Strings

For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Constraints:
1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.


Example 1:
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:
Input: str1 = "LEET", str2 = "CODE"
Output: ""

*/

func main() {
	println("Start main")
	var word1, word2 string
	word1 = "ABCABC"
	word2 = "ABC"
	fmt.Printf("%s\n", gcdOfStrings(word1, word2))
}

func gcdOfStrings(str1 string, str2 string) string {
	shortStr := ""
	longStr := ""

	if len(str1) < len(str2) {
		shortStr = str1
		longStr = str2
	} else if len(str1) > len(str2) {
		shortStr = str2
		longStr = str1
	} else {
		if str1 == str2 {
			return str1
		} else {
			return ""
		}
	}

	divider := ""

	divider = findDivider(longStr, shortStr)
	if divider != "" {
		return divider
	}

	for b := len(shortStr); b > 0; b-- {
		divider = findDivider(shortStr[b:], shortStr[:b])
		if divider != "" {

			divider = findDivider(longStr, divider)
			if divider != "" {
				break
			}
		}
	}

	return divider
}

func findDivider(str string, substr string) string {

	if len(str)%len(substr) != 0 {
		return ""
	}

	for i := 0; i+len(substr) <= len(str); i += len(substr) {

		for j := 0; j < len(substr); j++ {
			if substr[j] != str[i+j] {
				return ""
			}
		}
	}
	return substr
}
