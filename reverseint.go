package main

import (
	"strconv"
)

func reverseInt(n int) int {
	reversed := reverse(strconv.Itoa(n))
	i, _ := strconv.Atoi(reversed)
	return i
}
