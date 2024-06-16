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
			if i == 1 || i == y { // First and last row
				for j := 1; j <= x; j++ {
					if j == 1 || j == x { // First and last column
						fmt.Print("A")
					} else { // Middle columns
						fmt.Print("B")
					}
				}
				fmt.Println()
			} else { // Middle rows
				for j := 1; j <= x; j++ {
					if j == 1 || j == x { // First and last column
						fmt.Print("B")
					} else { // Middle columns
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
	}
}
