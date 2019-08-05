package main

func reversePrune(str string) string {
	rev := []rune(str)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

func reverse(str string) string {
	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}
