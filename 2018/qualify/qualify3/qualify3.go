package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

// var stderr *log.Logger

func main() {
	// // stderr = log.New(os.Stderr, "", 0)
	reader = bufio.NewReader(os.Stdin)
	strNum, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))

	for i := 0; i < num; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		A, _ := strconv.Atoi(line)
		ret := qualify3(A)
		if !ret {
			// stderr.Println("Fail")
			break
		} else {
			// stderr.Println("Sucess")
		}
	}
}

// 25 = 5 * 5
// 225 = 15 * 15
func qualify3(A int) bool {
	targetWidth := findTargetWidth(A)
	orchard := make([][]bool, targetWidth)
	for i := range orchard {
		orchard[i] = make([]bool, targetWidth)
	}

	x := 1
	y := 1
	round := 0
	for x != 0 && y != 0 {
		// stderr.Println("\n=====================")
		fillX, fillY := getNextFill(orchard)
		// stderr.Printf("Set %d %d\n", fillX+1, fillY+1)
		fmt.Printf("%d %d\n", fillX+1, fillY+1)
		x, y = readTwoInt()
		if x == -1 || y == -1 {
			// stderr.Printf("Error")
			return false
		}
		if x == 0 || y == 0 {
			break
		}
		round++
		// stderr.Printf("Get %d %d, round %d\n\n", x, y, round)
		orchard[x-1][y-1] = true
		printOrchard(orchard)
	}
	return true
}

func printOrchard(orchard [][]bool) {
	rowStr := "   "
	for idx := range orchard[0] {
		rowStr += fmt.Sprintf(" %d", idx+1)
	}
	// stderr.Println(rowStr)

	for rowIdx, row := range orchard {
		rowStr := fmt.Sprintf("%d: ", rowIdx+1)
		for _, cell := range row {
			if cell {
				rowStr += " o"
			} else {
				rowStr += "  "
			}
		}
		// stderr.Println(rowStr)
	}
}

func getNextFill(fill [][]bool) (x int, y int) {
	for i := 1; i < len(fill)-1; i++ {
		j := 1
		for ; j < len(fill[i])-1; j++ {
			if !fill[i-1][j-1] {
				return i, j
			}
			if i == len(fill)-2 {
				if !fill[i][j-1] || !fill[i+1][j-1] {
					return i, j
				}
			}

			if j == len(fill[i])-2 {
				if !fill[i-1][j] || !fill[i-1][j+1] {
					return i, j
				}
				if i == len(fill)-2 {
					if !fill[i][j] || !fill[i+1][j] || !fill[i][j+1] || !fill[i+1][j+1] {
						return i, j
					}
				}
			}
		}

		if i == len(fill)-2 {
			fmt.Println("WTF, all filled")
			break
		}
		// stderr.Printf("Finish row %d\n", i)
	}
	return 0, 0
}

func readTwoInt() (x int, y int) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	params := strings.Split(line, " ")
	x, _ = strconv.Atoi(params[0])
	y, _ = strconv.Atoi(params[1])
	return
}

func findTargetWidth(A int) int {
	for i := 1; ; i++ {
		if i*i > A {
			return i
		}
	}
}
