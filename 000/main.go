package main

import (
	"fmt"
	"log"
)

// 13. Roman to Integer

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
Given a roman numeral, convert it to an integer.
*/

func main() {

	println("XLIX", romanToInt("XLIX"))
	println("III", romanToInt("III"))
	println("LVIII", romanToInt("LVIII"))
	println("MCMXCIV", romanToInt("MCMXCIV"))
	println("CCXLVI", romanToInt("CCXLVI"))
	println("MMMDCCCLXXXVIII", romanToInt("MMMDCCCLXXXVIII"))
}

func xromanToInt(s string) int {
	/*const nI int = 1
	const nV int = 5
	const nX int = 10
	const nL int = 50
	const nC int = 100
	const nD int = 500
	const nM int = 1000

	const cI string = "I"
	const cV string = "V"
	const cX string = "X"
	const cL string = "L"
	const cC string = "C"
	const cD string = "D"
	const cM string = "M"*/
	/*const romans [7]string = [7]string{"I", "V", "X", "L", "C", "D", "M"}
	const arabics [7]int = [7]int{1, 5, 10, 50, 100, 500, 1000}*/

	var dict map[rune]int = map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result int
	//var prefix, postfix, result int = 0, 0,0

	// validate string
	// check for empty string, valid symbols and symbols order
	{
		//println("--- Validation ---")

		if len(s) == 0 {
			log.Fatal("inputStringValidationError:EmptyString")
		}

		var symbolRunePrev rune
		var symbolValuePrev int
		for symbolIdx, symbolRune := range s {
			//var symbolIdxPrev int
			symbolValue, exists := dict[symbolRune]
			//	fmt.Printf("Debug: index %d RunePrev: %c RuneCur: %c ValPrev: %d ValCur: %d\n", symbolIdx, symbolRunePrev, symbolRune, symbolValuePrev, symbolValue)
			// check for other symbols
			if !exists {
				log.Fatal("inputStringValidationError:InvalidSymbol")
			} else
			// check for order
			if symbolIdx > 0 && symbolValuePrev < symbolValue {
				/*
					I can be placed before V (5) and X (10) to make 4 and 9.
					X can be placed before L (50) and C (100) to make 40 and 90.
					C can be placed before D (500) and M (1000) to make 400 and 900.
				*/
				if !((symbolRunePrev == 'I' && (symbolRune == 'V' || symbolRune == 'X')) || (symbolRunePrev == 'X' && (symbolRune == 'L' || symbolRune == 'C')) || (symbolRunePrev == 'C' && (symbolRune == 'D' || symbolRune == 'M'))) {
					log.Fatal("inputStringValidationError:InvalidOrder")
				}
				/*
					check for IIII, XXXX, MMMM and LL, DD
				*/
			}
			//symbolIdxPrev, symbolRunePrev = symbolIdx, symbolRune
			symbolValuePrev = symbolValue
			symbolRunePrev = symbolRune
		}
	}

	// calculating result
	{
		println("--- Calculation ---")
		println(s)

		var preSum int = 0
		var symbolRunePrev rune
		var symbolValuePrev int
		for symbolIdx, symbolRune := range s {
			symbolValue := dict[symbolRune]

			if symbolIdx > 0 && symbolValuePrev < symbolValue {
				result += symbolValue - preSum
				preSum = 0
			} else if symbolIdx > 0 && symbolValuePrev > symbolValue {
				result += preSum
				preSum = symbolValue
			} else {
				preSum += symbolValue
			}
			fmt.Printf("Debug: index %d RunePrev: %c RuneCur: %c ValPrev: %d ValCur: %d preSum %d result %d\n", symbolIdx, symbolRunePrev, symbolRune, symbolValuePrev, symbolValue, preSum, result)

			symbolValuePrev = symbolValue
			symbolRunePrev = symbolRune
		}
		result += preSum
	}
	println(s, result)
	return result
}

func romansToInt(s string) int {
	/*code from chatGpt
	samesame but different*/
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int

	for i := 0; i < len(s); i++ {
		if i > 0 && values[s[i]] > values[s[i-1]] {
			result += values[s[i]] - 2*values[s[i-1]]
		} else {
			result += values[s[i]]
		}
	}
	println(s, result)
	return result
}

func romanToInt(s string) int {
	// submited issue

	var dict map[rune]int = map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result int
	{
		var preSum int = 0
		var symbolValuePrev int
		for symbolIdx, symbolRune := range s {
			symbolValue := dict[symbolRune]

			if symbolIdx > 0 && symbolValuePrev < symbolValue {
				result += symbolValue - preSum
				preSum = 0
			} else if symbolIdx > 0 && symbolValuePrev > symbolValue {
				result += preSum
				preSum = symbolValue
			} else {
				preSum += symbolValue
			}
			symbolValuePrev = symbolValue
		}
		result += preSum
	}
	return result
}
