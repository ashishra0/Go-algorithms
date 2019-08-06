package main

import "strings"

func capitalize(str string) string {
	words := []string{}
	newStr := strings.Split(str, " ")
	for _, word := range newStr {
		words = append(words, strings.ToUpper(word[:1])+word[1:])
	}
	return strings.Join(words, " ")
}
