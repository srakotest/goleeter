package main

import (
	"fmt"
	"strings"
)

/*
345. Reverse Vowels of a String

Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"

Constraints:
1 <= s.length <= 3 * 10^5
s consist of printable ASCII characters.
*/

func main() {
	println("Start main")
	fmt.Printf("%s\n", reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	var result strings.Builder
	j := len(s)
	var dict map[byte]int = map[byte]int{
		'a': 1, 'A': 1, 'e': 1, 'E': 1, 'i': 1, 'I': 1, 'o': 1, 'O': 1, 'u': 1, 'U': 1}

	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[i]]; ok {
			//vowel
			for j >= 0 {
				j--
				if _, ok := dict[s[j]]; ok {
					//vowel
					result.WriteByte(s[j])
					break
				}
			}
		} else {
			//consonant
			result.WriteByte(s[i])
		}
	}

	return result.String()
}
