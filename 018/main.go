package main

import (
	"fmt"
)

/*
392. Is Subsequence

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting
some (can be none) of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Constraints:
0 <= s.length <= 100
0 <= t.length <= 10^4
s and t consist only of lowercase English letters.



Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false

Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?

*/

func main() {
	println("Start main")

	s := "ace"
	t := "arcade"
	fmt.Printf("%v\n", isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen > tLen {
		return false
	}
	if sLen == tLen {
		if s == t {
			return true
		} else {
			return false
		}
	}
	if t == "" {
		return false
	}
	if s == "" {
		return true
	}

	k := 0 // counter for substring

	for _, v := range t {
		if byte(v) == s[k] {
			if k == sLen-1 {
				return true
			} else {
				k++
			}
		}
	}
	return false
}
