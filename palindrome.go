package main

func palindrome(str string) bool {
	reverseStr := reverse(str)
	if str == reverseStr {
		return true
	}
	return false
}
