package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	xStr := os.Args[1]
	yStr := os.Args[2]

	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)

	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						fmt.Print("A")
					} else if j == x {
						fmt.Print("C")
					} else {
						fmt.Print("B")
					}
				}
				fmt.Println()
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						fmt.Print("C")
					} else if j == x {
						fmt.Print("A")
					} else {
						fmt.Print("B")
					}
				}
				fmt.Println()
			} else {
				for j := 1; j <= x; j++ {
					if j == x || j == 1 {
						fmt.Print("B")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
	}
}
