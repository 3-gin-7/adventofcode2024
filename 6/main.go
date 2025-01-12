package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// assuming that start direction is always up and not at the end of the matrix
	direction := "up"
	hasNext := true
	count := 0
	debug := 0
	// run while loop until either i or j > len
	// run count and increase it everytime matrix[i][j] == "." and change it to "X"
	// determine the rotation mechanism.
	// direction string
	// update with switch case

	matrix, x, y := readFile()
	fmt.Println("working")

	for {
		if !hasNext {
			break
		}

		// inc count if cell is new and mark it with X
		if matrix[x][y] == "X" {
			// count++
		} else {
			matrix[x][y] = "X"
		}

		for {
			debug++
			if debug == 5779 {
				fmt.Println("hey")
			}
			// get next coords
			new_x, new_y := getNext(direction, x, y)
			// check out of bounds
			if new_x < 0 || new_y < 0 || new_x >= len(matrix) || new_y >= len(matrix[x]) {
				// out of bounds
				hasNext = false
				break
			} else {
				// not out of bounds
				if matrix[new_x][new_y] == "#" {
					direction = updateDirection(direction)
				} else {
					if matrix[new_x][new_y] != "X" {
						count++
					}
					hasNext = true
					x = new_x
					y = new_y
					break
				}
			}
		}
	}
	count++
	fmt.Printf("count is: %v", count)
}

func updateDirection(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "down":
		return "left"
	case "left":
		return "up"
	case "right":
		return "down"
	default:
		panic("unknown direction")
	}
}

func getNext(direction string, x int, y int) (int, int) {
	switch direction {
	case "up":
		x--
	case "down":
		x++
	case "left":
		y--
	case "right":
		y++
	default:
		panic("unknown direction")
	}

	return x, y
}

func readFile() ([][]string, int, int) {
	output := [][]string{}
	x, y, count := 0, 0, 0
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)

	if err != nil {
		fmt.Printf("Failed to open file with error:%v\r\n", err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)

	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, "^") {
			x = count
			y = strings.Index(line, "^")
		}
		output = append(output, strings.Split(line, ""))
		count++
	}

	return output, x, y
}
