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
	count := 0

	for i, line := range matrix {
		// search for x and s

		// fmt.Println(line)
		// fmt.Println(i)
		for j, char := range line {
			// check if the char is beginning for 'xmas' or 'samx'
			if char != "X" {
				continue
			}

			// forward search (j+4)
			if j+3 < len(line) {
				sub := line[j : j+4]
				if checkSubstring(strings.Join(sub, "")) {
					count++
				}
			}
			// backward search (j-4)
			if j-3 > 0 {
				sub := line[j-3 : j+1]
				if checkSubstring(strings.Join(sub, "")) {
					count++
				}
			}
			// upward search (i+4)
			if i-3 >= 0 {
				sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i-1][j], matrix[i-2][j], matrix[i-3][j])
				if checkSubstring(sub) {
					count++
				}
			}
			// downward search (i-4)
			if i+3 < len(matrix) {
				sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j])
				if checkSubstring(sub) {
					count++
				}
			}
			// NE search (i+4, j+4)
			if i-3 >= 0 && j+3 < len(line) {
				sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i-1][j+1], matrix[i-2][j+2], matrix[i-3][j+3])
				if checkSubstring(sub) {
					count++
				}
			}
			// SE search (i-4, j+4)
			if i+3 < len(matrix) && j+3 < len(line) {
				sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3])
				if checkSubstring(sub) {
					count++
				}
			}
			// SW search (i-4, j-4)
			if i+3 < len(matrix) && j-3 >= 0 {
				sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i+1][j-1], matrix[i+2][j-2], matrix[i+3][j-3])
				if checkSubstring(sub) {
					count++
				}
			}
			// NW search (i+4, j-4)
			if i-3 >= 0 && j-3 >= 0 {
				sub := fmt.Sprintf("%v%v%v%v", matrix[i][j], matrix[i-1][j-1], matrix[i-2][j-2], matrix[i-3][j-3])
				if checkSubstring(sub) {
					count++
				}
			}
			// fmt.Println(j)
			// fmt.Println(char)
		}
		// fmt.Println("hey")
	}

	fmt.Printf("Count is: %v", count)
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
