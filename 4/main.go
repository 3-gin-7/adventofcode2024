package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("working")
	matrix := readInput()
	part_one_count := 0
	part_two_count := 0

	for i, line := range matrix {
		for j, char := range line {
			if char == "X" {
				// count all of the xmas and samx instances
				part_one_count += countXmas(i, j, line, matrix)
			} else if char == "A" {
				// counts crosses where mas or sam intersect at a
				// apply backwards search as well
				part_two_count += countCrossMas(i, j, line, matrix)
			} else {
				continue
			}
		}
	}

	fmt.Printf("Part part one: %v\r\n", part_one_count)
	fmt.Printf("Part part two: %v\r\n", part_two_count)
}

func countXmas(i int, j int, line []string, matrix [][]string) int {
	part_one_count := 0

	// forward search (j+4)
	if j+3 < len(line) {
		sub := line[j : j+4]
		if checkSubstring(strings.Join(sub, "")) {
			part_one_count++
		}
	}
	// backward search (j-4)
	if j-3 > 0 {
		sub := line[j-3 : j+1]
		if checkSubstring(strings.Join(sub, "")) {
			part_one_count++
		}
	}
	// upward search (i+4)
	if i-3 >= 0 {
		sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i-1][j], matrix[i-2][j], matrix[i-3][j])
		if checkSubstring(sub) {
			part_one_count++
		}
	}
	// downward search (i-4)
	if i+3 < len(matrix) {
		sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j])
		if checkSubstring(sub) {
			part_one_count++
		}
	}
	// NE search (i+4, j+4)
	if i-3 >= 0 && j+3 < len(line) {
		sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i-1][j+1], matrix[i-2][j+2], matrix[i-3][j+3])
		if checkSubstring(sub) {
			part_one_count++
		}
	}
	// SE search (i-4, j+4)
	if i+3 < len(matrix) && j+3 < len(line) {
		sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3])
		if checkSubstring(sub) {
			part_one_count++
		}
	}
	// SW search (i-4, j-4)
	if i+3 < len(matrix) && j-3 >= 0 {
		sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i+1][j-1], matrix[i+2][j-2], matrix[i+3][j-3])
		if checkSubstring(sub) {
			part_one_count++
		}
	}
	// NW search (i+4, j-4)
	if i-3 >= 0 && j-3 >= 0 {
		sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i-1][j-1], matrix[i-2][j-2], matrix[i-3][j-3])
		if checkSubstring(sub) {
			part_one_count++
		}
	}

	return part_one_count
}

func countCrossMas(i int, j int, line []string, matrix [][]string) int {
	part_two_count := 0

	// search
	if i-1 >= 0 && i+1 < len(matrix) && j-1 >= 0 && j+1 < len(line) {
		first_diagonal := fmt.Sprintf("%v%v%v", matrix[i-1][j-1], matrix[i][j], matrix[i+1][j+1])
		second_diagonal := fmt.Sprintf("%v%v%v", matrix[i-1][j+1], matrix[i][j], matrix[i+1][j-1])
		if (first_diagonal == "MAS" || first_diagonal == "SAM") && (second_diagonal == "MAS" || second_diagonal == "SAM") {
			part_two_count++
		}
	}
	return part_two_count
}

func readInput() [][]string {
	// file name
	FILE_NAME := "input.txt"
	// FILE_NAME := "test.txt"
	output := [][]string{}

	fi, err := os.Open(FILE_NAME)

	if err == nil {
		defer fi.Close()
	}

	if err != nil {
		fmt.Printf("Failed to open file {%v} with error: %v", FILE_NAME, err)
		return nil
	}

	sc := bufio.NewScanner(fi)

	for sc.Scan() {
		line := sc.Text()
		arr := strings.Split(line, "")
		output = append(output, arr)
	}

	return output
}

func checkSubstring(sub string) bool {
	XMAS := "XMAS"
	SAMX := "SAMX"
	return sub == XMAS || sub == SAMX
}
