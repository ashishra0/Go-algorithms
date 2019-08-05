package main

import (
	"strings"
)

func vowelCount(str string) int {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0
	for i := 0; i < len(vowels); i++ {
		if strings.Contains(str, vowels[i]) {
			count++
		}
	}
	return count
}
