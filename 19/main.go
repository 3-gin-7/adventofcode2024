package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	p, d := readFile(false)

	// fmt.Printf("patterns: %v\r\n", p)
	// fmt.Printf("designs: %v\r\n", d)

	part_one_count := getPartOneCount(p, d)
	fmt.Printf("Part One Count: %d\r\n", part_one_count)
}

// dynamic programming - minimum word-break
func getPartOneCount(patterns, designs []string) int {
	count := 0

	for _, d := range designs {
		n := len(d)
		dp := make([]int, n+1)

		// init to infinity
		for i, _ := range dp {
			dp[i] = math.MaxInt
		}
		dp[0] = 0

		for i := 0; i < n; i++ {
			if dp[i] == math.MaxInt {
				continue
			}

			for _, p := range patterns {
				p_len := len(p)
				if d[i:i+p_len] == p {
					dp[i+p_len] = min(dp[i+p_len], dp[i]+1)
				}
			}
		}

		// check if design can be made
		if dp[n] != math.MaxInt {
			count++
		}
	}

	return count
}

func readFile(is_test bool) ([]string, []string) {
	FILE_NAME := ""
	var patterns []string
	var designs []string

	if is_test {
		FILE_NAME = "test.txt"
	} else {
		FILE_NAME = "input.txt"
	}

	file, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}, []string{}
	}
	defer file.Close()

	is_pattern := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			is_pattern = false
			continue
		}

		if is_pattern {
			split := strings.SplitSeq(scanner.Text(), ",")
			for s := range split {
				patterns = append(patterns, strings.Trim(s, " "))
			}
		} else {
			designs = append(designs, scanner.Text())
		}
	}

	return patterns, designs
}
