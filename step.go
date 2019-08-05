package main

import "fmt"

func getSteps(n int) {
	for row := 0; row < n; row++ {
		stair := ""
		for column := 0; column < n; column++ {
			if column <= row {
				stair += "*"
			} else {
				stair += " "
			}
		}
		fmt.Println(stair)
	}
}
