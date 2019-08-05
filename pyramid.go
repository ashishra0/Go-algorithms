package main

import "fmt"

func getPyramid(n int) {
	midpoint := (2*n - 1) / 2
	for row := 0; row < n; row++ {
		level := ""
		for column := 0; column < 2*n-1; column++ {
			if midpoint-row <= column && midpoint+row >= column {
				level += "*"
			} else {
				level += " "
			}
		}
		fmt.Println(level)
	}
}
