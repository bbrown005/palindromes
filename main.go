package main

import "fmt"

func IsPalindrome(testString string) bool {
	// set left to 0
	left := 0
	// set right to last index of string
	right := len(testString) - 1
	// while left is before right
	for left < right {

		if testString[left] != testString[right] {
			return false
		}
		// move left one right
		left++
		// move right one left
		right--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("racecar"))
}
