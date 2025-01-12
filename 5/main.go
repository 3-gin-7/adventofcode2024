package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	broken_sum := 0
	broken_lines := []string{}

	matrix, lines := readFile()

	// fmt.Println(matrix)
	// fmt.Println(lines)

	// loop over the list of inputs
	for j, input_list := range lines {
		fmt.Printf("Checking line:%v\r\n", j)
		isValid := true
		pages := strings.Split(input_list, ",")
		for i, page_num := range pages {
			remain := pages[i+1:]
			isValid = lookupVal(page_num, remain, matrix)
			if !isValid {
				// not valid
				// take the line and pass to fix method and get the middle
				broken_lines = append(broken_lines, input_list)
				// broken_sum += fixOrderAndGetMiddle(i, pages, matrix)
				break
			}
		}

		if isValid {
			fmt.Printf("{%v} is valid\r\n", j)
			sum += getMiddle(pages)
		} else {
			fmt.Printf("{%v} is not valid\r\n", j)
		}
	}

	fmt.Println("Broken lines")
	for _, i := range broken_lines {
		fixed_line := fixOrder(strings.Split(i, ","), matrix)
		broken_sum += getMiddle(fixed_line)
	}

	fmt.Printf("\r\n\r\nThe sum is: %v\r\n", sum)
	fmt.Printf("\r\n\r\nThe broken sum is: %v\r\n", broken_sum)
}

func lookupVal(val string, items []string, matrix map[string][]string) bool {
	for _, i := range items {
		value, ok := matrix[i]
		if !ok {
			return true
		}

		for _, num := range value {
			if num == val {
				return false
			}
		}
	}
	return true
}

func fixOrder(line []string, matrix map[string][]string) []string {
	swapped := false
	for val := 0; val < len(line); val++ {
		for i := val + 1; i < len(line); i++ {
			allowed, ok := matrix[line[i]]
			if !ok {
				continue
			}

			for _, r := range allowed {
				if r == line[val] {
					tmp := line[val]
					line[val] = line[i]
					line[i] = tmp
					swapped = true
				}
			}
		}
		if swapped {
			swapped = false
			val--
		}
	}

	return line
}

func getMiddle(items []string) int {
	// fetch middle value
	if len(items) == 1 {
		i, _ := strconv.Atoi(items[0])
		return i
	}
	if len(items)%2 != 1 {
		fmt.Printf("no middle in the array: %v", items)
		return 0
	} else {
		index := (len(items) + 1) / 2
		i, _ := strconv.Atoi(items[index-1])
		return i
	}
}

func readFile() (map[string][]string, []string) {
	output := make(map[string][]string)
	output2 := []string{}
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"
	flag := false

	fi, err := os.Open(FILE_NAME)

	if err == nil {
		defer fi.Close()
	}

	sc := bufio.NewScanner(fi)

	for sc.Scan() {
		line := sc.Text()
		if !flag && line != "" {
			nums := strings.Split(line, "|")
			output[nums[0]] = append(output[nums[0]], nums[1])
		} else if flag {
			output2 = append(output2, line)
		}

		if line == "" {
			flag = true
		}
	}

	return output, output2
}
