package main

import (
	"fmt"
	"strings"
)

/*
1768. Merge Strings Alternately

You are given two strings word1 and word2. Merge the strings by adding letters in alternating order,
starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
Return the merged string.

Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.

Example 1:
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r

Example 2:
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s

Example 3:
Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d

*/

func main() {
	println("Start main")
	var word1, word2 string
	word1 = "abc"
	word2 = "pqf"
	fmt.Printf("%s\n", mergeAlternately(word1, word2))
	fmt.Printf("%s\n", mergeAlternatelySlices(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	var result strings.Builder
	result.Grow(len(word1) + len(word2))
	i, j := 0, 0

	for i+j < len(word1)+len(word2) {
		if i < len(word1) {
			result.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			result.WriteByte(word2[j])
			j++
		}
	}
	return result.String()
}

func mergeAlternatelySlices(word1 string, word2 string) string {
	var result strings.Builder
	result.Grow(len(word1) + len(word2))
	minLen := 0
	remainStr := ""

	if len(word1) <= len(word2) {
		minLen = len(word1)
		remainStr = word2[minLen:]
	} else {
		minLen = len(word2)
		remainStr = word1[minLen:]
	}

	for i := 0; i < minLen; i++ {
		/*result.WriteByte(word1[i])
		result.WriteByte(word2[i])*/
		result.Write([]byte{word1[i], word2[i]})
	}
	result.WriteString(remainStr)

	return result.String()
}
