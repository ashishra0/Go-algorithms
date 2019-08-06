package main

import (
	"sort"
	"strings"
)

func isAnagram(strA, strB string) bool {
	// result := []string{}
	// for _, candidate := range candidates {
	// 	if normalize(subject) == normalize(candidate) &&
	// 		strings.ToLower(subject) != strings.ToLower(candidate) {
	// 		result = append(result, strings.ToLower(candidate))
	// 	}
	// }
	// return result
	a := normalize(strA)
	b := normalize(strB)

	if a == b {
		return true
	}
	return false
}

func normalize(subject string) string {
	runes := []string{}
	for _, Rune := range subject {
		runes = append(runes, strings.ToLower(string(Rune)))
	}
	sort.Strings(runes)
	return strings.Join(runes, "")
}
