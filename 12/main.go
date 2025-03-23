package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cache := make(map[string]bool)
	type_count := make(map[string]int)

	matrix := readFile()
	count := 0
	corner_count := 0

	for i, line := range matrix {
		for j, val := range line {
			perimeter := 0
			corners := 0
			key := fmt.Sprintf("%v,%v", i, j)
			_, ok := cache[key]
			if !ok {
				perimeter, corners = getCountTwo(matrix, cache, type_count, key, val)
				count += perimeter * type_count[val]
				corner_count += corners * type_count[val]
			}
			type_count[val] = 0
		}
	}

	fmt.Printf("Count is: %v\r\n", count)
	fmt.Printf("Corner count is: %v\r\n", corner_count)
}

func getCountTwo(matrix [][]string, cache map[string]bool, type_count map[string]int, key, val string) (int, int) {
	count := 0
	corner_count := 0

	split := strings.Split(key, ",")
	i, _ := strconv.Atoi(split[0])
	j, _ := strconv.Atoi(split[1])

	// base case
	if i < 0 || j < 0 || i > len(matrix)-1 || j > len(matrix[i])-1 || matrix[i][j] != val {
		count += 1
		return count, corner_count
	}

	_, ok := cache[key]

	if !ok {
		cache[key] = true
		type_count[val]++
	} else {
		return count, corner_count
	}

	up := fmt.Sprintf("%v,%v", i-1, j)
	down := fmt.Sprintf("%v,%v", i+1, j)
	left := fmt.Sprintf("%v,%v", i, j-1)
	right := fmt.Sprintf("%v,%v", i, j+1)

	// check every side for != val
	next_up := getFromMatrix(i-1, j, matrix, val)
	next_down := getFromMatrix(i+1, j, matrix, val)
	next_right := getFromMatrix(i, j+1, matrix, val)
	next_left := getFromMatrix(i, j-1, matrix, val)

	matches := getBoolCount(next_up, next_down, next_right, next_left)

	if matches == 0 {
		corner_count += 4
	} else if matches == 1 {
		corner_count += 2
	} else if matches == 2 && ((next_up && next_down) || (next_right && next_left)) {
		corner_count += 0
	} else {
		if next_left && next_up {
			if matches == 2 && matrix[i-1][j-1] != val {
				corner_count += 2
			} else if matches > 2 && matrix[i-1][j-1] == val {
				corner_count += 0
			} else {
				corner_count += 1
			}
		}

		if next_right && next_up {
			if matches == 2 && matrix[i-1][j+1] != val {
				corner_count += 2
			} else if matches > 2 && matrix[i-1][j+1] == val {
				corner_count += 0
			} else {
				corner_count += 1
			}
		}

		if next_right && next_down {
			if matches == 2 && matrix[i+1][j+1] != val {
				corner_count += 2
			} else if matches > 2 && matrix[i+1][j+1] == val {
				corner_count += 0
			} else {
				corner_count += 1
			}
		}
		if next_left && next_down {
			if matches == 2 && matrix[i+1][j-1] != val {
				corner_count += 2
			} else if matches > 2 && matrix[i+1][j-1] == val {
				corner_count += 0
			} else {
				corner_count += 1
			}
		}
	}

	inc_count, inc_corner_count := getCountTwo(matrix, cache, type_count, up, val)
	count += inc_count
	corner_count += inc_corner_count
	inc_count, inc_corner_count = getCountTwo(matrix, cache, type_count, right, val)
	count += inc_count
	corner_count += inc_corner_count
	inc_count, inc_corner_count = getCountTwo(matrix, cache, type_count, down, val)
	count += inc_count
	corner_count += inc_corner_count
	inc_count, inc_corner_count = getCountTwo(matrix, cache, type_count, left, val)
	count += inc_count
	corner_count += inc_corner_count

	return count, corner_count
}

func getBoolCount(next_up, next_right, next_left, next_down bool) int {
	output := 0
	if next_up {
		output += 1
	}

	if next_right {
		output += 1
	}

	if next_left {
		output += 1
	}

	if next_down {
		output += 1
	}
	return output
}

func getFromMatrix(i, j int, matrix [][]string, val string) bool {
	if i < 0 || i > len(matrix)-1 {
		return false
	}

	if j < 0 || j > len(matrix[i])-1 {
		return false
	}

	return matrix[i][j] == val
}

func readFile() [][]string {
	// FILE_NAME := "test.txt"
	// FILE_NAME := "hey.txt"
	FILE_NAME := "input.txt"

	output := [][]string{}

	fi, err := os.Open(FILE_NAME)

	if err != nil {
		panic(err)
	} else {
		defer fi.Close()
	}

	sc := bufio.NewScanner(fi)

	for sc.Scan() {
		line := sc.Text()
		output = append(output, strings.Split(line, ""))
	}

	return output
}
