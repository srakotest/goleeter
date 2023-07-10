package main

import (
	"fmt"
	"strconv"
)

/*
412. Fizz Buzz

Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.

Example 1:
Input: n = 3
Output: ["1","2","Fizz"]

Example 2:
Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]

Example 3:
Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
*/
func main() {
	println("Start main")
	println(fmt.Sprintln(fizzBuzz(10)))
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	var isFizz, isBuzz bool
	for i := 1; i <= n; i++ {
		isFizz = ((i % 3) == 0)
		isBuzz = ((i % 5) == 0)
		switch {
		case isFizz && isBuzz:
			result[i-1] = "FizzBuzz"
		case isFizz:
			result[i-1] = "Fizz"
		case isBuzz:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
