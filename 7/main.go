package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ops_permut := make(map[string][][]string)
	result_sum := 0

	results, factors, max_factors := readFile()

	// generate map with that contains all the permutations
	for i := 1; i <= max_factors; i++ {
		getOpsPermutes(make([]string, i), i-1, 0, &ops_permut, strconv.Itoa(i))
	}

	// run though the results, factors and permutations
	for i := 0; i < len(results); i++ {
		int_res, err := strconv.Atoi(results[i])
		if err != nil {
			panic(err)
		}

		factor := strings.Split(factors[i], " ")
		int_factors := []int{}
		for _, i := range factor {
			conv, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			int_factors = append(int_factors, conv)
		}

		permute_count := len(int_factors) - 1
		test := checkResults(int_res, int_factors, ops_permut[strconv.Itoa(permute_count)])

		if test {
			// fmt.Printf("Success for result: %v\r\n", int_res)
			result_sum += int_res
		}
	}

	fmt.Printf("Part one results sum: %v\r\n", result_sum)
}

func checkResults(result int, factors []int, ops [][]string) bool {
	for _, i := range ops {
		end_res := factors[0]

		for count, j := range i {
			if j == "+" {
				end_res += factors[count+1]
			} else if j == "*" {
				end_res *= factors[count+1]
			}
		}

		if end_res == result {
			return true
		}
	}

	return false
}

func getOpsPermutes(data []string, last int, index int, ops_permut *map[string][][]string, loop string) {
	ops := []string{"+", "*"}
	for i := 0; i < len(ops); i++ {

		data[index] = ops[i]

		if index == last {
			// fmt.Println(data)
			addOutputToMap(ops_permut, loop, strings.Join(data, ":"))
		} else {
			getOpsPermutes(data, last, index+1, ops_permut, loop)
		}
	}
}

func addOutputToMap(ops_permut *map[string][][]string, index string, input string) {
	(*ops_permut)[index] = append((*ops_permut)[index], strings.Split(input, ":"))
}

func readFile() ([]string, []string, int) {
	results := []string{}
	factors := []string{}
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)

	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)

	max_factors := 0

	for sc.Scan() {
		line := sc.Text()
		split := strings.Split(line, ":")
		results = append(results, split[0])
		trimmed_factors := strings.Trim(split[1], " ")

		factors = append(factors, trimmed_factors)

		if len(strings.Split(trimmed_factors, " ")) > max_factors {
			max_factors = len(strings.Split(trimmed_factors, " "))
		}
	}

	max_factors--

	return results, factors, max_factors
}
