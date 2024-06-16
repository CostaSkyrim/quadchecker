package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Check command-line arguments
	if len(os.Args) == 1 {
		// No arguments provided, run quadchecker functionality
		if err := runQuadchecker(); err != nil {
			fmt.Println("Error running quadchecker:", err)
		}
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go build <binary>")
		return
	}

	action := os.Args[1]
	target := os.Args[2]

	if action != "build" {
		fmt.Println("Unknown action:", action)
		return
	}

	// Build the specified binary
	switch target {
	case "quadA":
		if err := writeAndBuild("quadA", quadAContent); err != nil {
			fmt.Println("Error building quadA binary:", err)
		} else {
			fmt.Println("Built quadA binary successfully!")
		}
	case "quadB":
		if err := writeAndBuild("quadB", quadBContent); err != nil {
			fmt.Println("Error building quadB binary:", err)
		} else {
			fmt.Println("Built quadB binary successfully!")
		}
	case "quadC":
		if err := writeAndBuild("quadC", quadCContent); err != nil {
			fmt.Println("Error building quadC binary:", err)
		} else {
			fmt.Println("Built quadC binary successfully!")
		}
	case "quadD":
		if err := writeAndBuild("quadD", quadDContent); err != nil {
			fmt.Println("Error building quadD binary:", err)
		} else {
			fmt.Println("Built quadD binary successfully!")
		}
	case "quadE":
		if err := writeAndBuild("quadE", quadEContent); err != nil {
			fmt.Println("Error building quadE binary:", err)
		} else {
			fmt.Println("Built quadE binary successfully!")
		}
	case "quadchecker":
		if err := writeAndBuild("quadchecker", quadcheckerContent); err != nil {
			fmt.Println("Error building quadchecker binary:", err)
		} else {
			fmt.Println("Built quadchecker binary successfully!")
		}
	default:
		fmt.Println("Unknown target:", target)
		return
	}

	// Clean up temporary .go files
	cleanup(target + ".go")
}

// Function to write source code to a file and build the binary
func writeAndBuild(binaryName, content string) error {
	fileName := binaryName + ".go"
	if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing %s: %w", fileName, err)
	}
	if err := buildBinary(binaryName, fileName); err != nil {
		return fmt.Errorf("error building %s binary: %w", binaryName, err)
	}
	return nil
}

// Function to build a binary file from a given .go file
func buildBinary(outputName, sourceFile string) error {
	cmd := exec.Command("go", "build", "-o", outputName, sourceFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Function to run the quadchecker functionality
func runQuadchecker() error {
	var arr []rune

	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		arr = append(arr, char)
	}
	str := string(arr)

	x := 0
	y := 0
	for _, char := range arr {
		if char != '\n' && y == 0 {
			x++
		}
		if char == '\n' {
			y++
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return nil
	}
	if isEqual(str, x, y, 'o', 'o', 'o', 'o', '-', '|') {
		fmt.Printf("[quadA] [%v] [%v]\n", x, y)
		return nil
	}

	if isEqual(str, x, y, '/', '\\', '\\', '/', '*', '*') {
		fmt.Printf("[quadB] [%v] [%v]\n", x, y)
		return nil
	}

	n := 0
	if isEqual(str, x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
		n++
		fmt.Printf("[quadC] [%v] [%v]", x, y)
	}
	if isEqual(str, x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[quadD] [%v] [%v]", x, y)
	}
	if isEqual(str, x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[quadE] [%v] [%v]", x, y)
	}

	if n > 0 {
		fmt.Println()
		return nil
	}

	fmt.Println("Not a Raid function")
	return nil
}

func isEqual(str string, x, y int, c1, c2, c3, c4, hor, ver rune) bool {
	var arrE []rune
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 {
				if j == 0 {
					arrE = append(arrE, c1)
				} else if j == x-1 {
					arrE = append(arrE, c2)
				} else {
					arrE = append(arrE, hor)
				}
			} else if i == y-1 {
				if j == 0 {
					arrE = append(arrE, c3)
				} else if j == x-1 {
					arrE = append(arrE, c4)
				} else {
					arrE = append(arrE, hor)
				}
			} else {
				if j == 0 || j == x-1 {
					arrE = append(arrE, ver)
				} else {
					arrE = append(arrE, ' ')
				}
			}
		}
		arrE = append(arrE, '\n')
	}
	strE := string(arrE)
	return strE == str
}

// Function to clean up temporary .go files
func cleanup(files ...string) {
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			fmt.Println("Error cleaning up", file, ":", err)
		}
	}
}

// Content for each source file
var (
	quadAContent = `package main

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
			if i == 1 || i == y {
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						fmt.Print("o")
					} else {
						fmt.Print("-")
					}
				}
				fmt.Println()
			} else {
				for j := 1; j <= x; j++ {
					if j == x || j == 1 {
						fmt.Print("|")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
	}
}
`
	quadBContent = `package main

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
						fmt.Print("/")
					} else if j == x {
						fmt.Print("\\")
					} else {
						fmt.Print("*")
					}
				}
				fmt.Println()
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						fmt.Print("\\")
					} else if j == x {
						fmt.Print("/")
					} else {
						fmt.Print("*")
						}
				}
				fmt.Println()
			} else {
				for j := 1; j <= x; j++ {
					if j == x || j == 1 {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
	}
}
`
	quadCContent = `package main

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
						fmt.Print("A")
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
						fmt.Print("C")
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
`
	quadDContent = `package main

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
						fmt.Print("A")
					} else if j == x {
						fmt.Print("C")
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
`
	quadEContent = `package main

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
`
	quadcheckerContent = `package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arr []rune

	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		arr = append(arr, char)
	}
	str := string(arr)

	x := 0
	y := 0
	for _, char := range arr {
		if char != '\n' && y == 0 {
			x++
		}
		if char == '\n' {
			y++
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	if isEqual(str, x, y, 'o', 'o', 'o', 'o', '-', '|') {
		fmt.Printf("[quadA] [%v] [%v]\n", x, y)
		return
	}

	if isEqual(str, x, y, '/', '\\', '\\', '/', '*', '*') {
		fmt.Printf("[quadB] [%v] [%v]\n", x, y)
		return
	}

	n := 0
	if isEqual(str, x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
		n++
		fmt.Printf("[quadC] [%v] [%v]", x, y)
	}
	if isEqual(str, x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[quadD] [%v] [%v]", x, y)
	}
	if isEqual(str, x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[quadE] [%v] [%v]", x, y)
	}

	if n > 0 {
		fmt.Println()
		return
	}

	fmt.Println("Not a Raid function")
}

func isEqual(str string, x, y int, c1, c2, c3, c4, hor, ver rune) bool {
	var arrE []rune
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 {
				if j == 0 {
					arrE = append(arrE, c1)
				} else if j == x-1 {
					arrE = append(arrE, c2)
				} else {
					arrE = append(arrE, hor)
				}
			} else if i == y-1 {
				if j == 0 {
					arrE = append(arrE, c3)
				} else if j == x-1 {
					arrE = append(arrE, c4)
				} else {
					arrE = append(arrE, hor)
				}
			} else {
				if j == 0 || j == x-1 {
					arrE = append(arrE, ver)
				} else {
					arrE = append(arrE, ' ')
				}
			}
		}
		arrE = append(arrE, '\n')
	}
	strE := string(arrE)
	if strE == str {
		return true
	}
	return false
}
`
)
