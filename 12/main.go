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

	// loop over the array
	// use recursion and save to seen dictionary
	matrix := readFile()
	count := 0

	for i, line := range matrix {
		for j, val := range line {
			perimeter := 0
			key := fmt.Sprintf("%v,%v", i, j)
			_, ok := cache[key]
			if !ok {
				perimeter = getCount(count, matrix, cache, type_count, key, val)
				count += perimeter * type_count[val]
			}
			type_count[val] = 0
			// fmt.Println(count)
		}
	}

	fmt.Printf("Count is: %v\r\n", count)
}

func getCount(count int, matrix [][]string, cache map[string]bool, type_count map[string]int, key, val string) int {
	count = 0
	split := strings.Split(key, ",")
	i, _ := strconv.Atoi(split[0])
	j, _ := strconv.Atoi(split[1])

	// base case
	if i < 0 || j < 0 || i > len(matrix)-1 || j > len(matrix)-1 || matrix[i][j] != val {
		count += 1
		return count
	}

	_, ok := cache[key]

	if !ok {
		cache[key] = true
		type_count[val]++
	} else {
		return count
	}

	count += getCount(count, matrix, cache, type_count, fmt.Sprintf("%v,%v", i-1, j), val)
	count += getCount(count, matrix, cache, type_count, fmt.Sprintf("%v,%v", i, j+1), val)
	count += getCount(count, matrix, cache, type_count, fmt.Sprintf("%v,%v", i+1, j), val)
	count += getCount(count, matrix, cache, type_count, fmt.Sprintf("%v,%v", i, j-1), val)

	return count
}

func readFile() [][]string {
	// FILE_NAME := "test.txt"
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
