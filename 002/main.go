package main

/*
383. Ransom Note

Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.

Constraints:
1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.



Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true

*/
func main() {
	println("Start main")
	println(canConstruct("aaba", "aasdasddsadddsadaaasdsasd"))
}

func canConstruct(ransomNote string, magazine string) bool {
	var alphabet [26]int
	//println("canConstruct")

	if len(ransomNote) > len(magazine) {
		return false
	}

	// fill alphabet
	for i := 0; i < len(magazine); i++ {
		alphabet[magazine[i]-'a']++
	}

	//print alphabet
	/*for i := 0; i < 26; i++ {
		fmt.Printf("%c-%d; ", i+'a', alphabet[i])
	}
	println()*/

	// check ransomNote
	for i := 0; i < len(ransomNote); i++ {
		if alphabet[ransomNote[i]-'a'] == 0 {
			return false
		} else {
			alphabet[ransomNote[i]-'a']--
		}
	}

	//print alphabet
	/*for i := 0; i < 26; i++ {
		fmt.Printf("%c-%d; ", i+'a', alphabet[i])
	}
	println()*/

	return true
}
