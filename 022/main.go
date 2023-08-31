package main

import (
	"fmt"
)

/*
1456. Maximum Number of Vowels in a Substring of Given Length

Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Constraints:
1 <= s.length <= 10^5
s consists of lowercase English letters.
1 <= k <= s.length

Example 1:
Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.

Example 2:
Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.

Example 3:
Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

*/

func main() {
	println("Start main")
	task := "leetcoder"
	fmt.Printf("\"%v\"\n", maxVowels(task, 3))
}

func isVowel(char byte) bool {
	switch char {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	default:
		return false
	}
}

func maxVowels(s string, k int) int {
	tmpResult := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			tmpResult++
		}
	}
	maxResult := tmpResult

	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			tmpResult++
		}
		if isVowel(s[i-k]) {
			tmpResult--
		}
		if tmpResult == k {
			return k
		}
		if tmpResult > maxResult {
			maxResult = tmpResult
		}
	}
	return maxResult

}
