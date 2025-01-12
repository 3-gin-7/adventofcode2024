package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// assuming that start direction is always up and not at the end of the matrix
	matrix, x, y := readFile()

	// part one
	part_one_count, _ := getStepsCount(&matrix, x, y, false)
	// part_one_count := 0

	// part two
	part_two_count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != "#" {
				tmp := matrix[i][j]
				matrix[i][j] = "#"
				_, hasCycle := getStepsCount(&matrix, x, y, true)

				if hasCycle {
					part_two_count++
				}

				matrix[i][j] = tmp
			}
		}
	}

	fmt.Printf("part one count is: %v\r\n", part_one_count)
	fmt.Printf("part two count is: %v\r\n", part_two_count)
}

func getStepsCount(matrix *[][]string, x int, y int, detectCycles bool) (int, bool) {
	count := 0
	direction := "up"
	hasNext := true
	hasCycle := false
	visitedMap := make(map[string][]string)

	// mark the initial position
	count++
	(*matrix)[x][y] = "X"

	for {
		if !hasNext {
			break
		}

		if detectCycles {
			key := fmt.Sprintf("%v|%v", x, y)
			value, ok := visitedMap[key]

			if ok && checkSliceForDirection(value, direction) {
				// same position with the same direction
				hasCycle = true
				break
			}

			// map the visited node
			visitedMap[key] = append(visitedMap[key], direction)
		}

		// get next coords
		new_x, new_y := getNext(direction, x, y)
		// check out of bounds
		if new_x < 0 || new_y < 0 || new_x >= len(*matrix) || new_y >= len((*matrix)[x]) {
			// out of bounds
			hasNext = false
		} else {
			// not out of bounds
			if (*matrix)[new_x][new_y] == "#" {
				direction = updateDirection(direction)
			} else {
				x = new_x
				y = new_y

				if (*matrix)[x][y] != "X" {
					count++
					(*matrix)[x][y] = "X"
				}
			}
		}
	}

	return count, hasCycle
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

func checkSliceForDirection(slice []string, direction string) bool {
	for _, i := range slice {
		if i == direction {
			return true
		}
	}

	return false
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
