package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	ops_permut := make(map[string][][]string)
	result_sum := 0
	max_failed_factor_len := 0
	failed_int_results := []int{}
	failed_factors := []string{}

	results, factors, max_factors := readFile()

	// generate map with that contains all the permutations
	for i := 1; i <= max_factors; i++ {
		getOpsPermutes(make([]string, i), i-1, 0, &ops_permut, strconv.Itoa(i), false)
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

		if checkResults(int_res, int_factors, ops_permut[strconv.Itoa(permute_count)]) {
			result_sum += int_res
		} else {
			// if result is false, add combination for the part two input
			// check if the len of factors == 2 and test the concat separately
			if len(int_factors)-1 > max_failed_factor_len {
				max_failed_factor_len = len(int_factors) - 1
			}
			failed_int_results = append(failed_int_results, int_res)
			failed_factors = append(failed_factors, factors[i])
		}
	}

	fmt.Printf("Part one results sum: %v\r\n", result_sum)

	// part two
	ops_permut = make(map[string][][]string, max_failed_factor_len)

	// generate ops permutations with concatenation
	for i := 1; i <= max_failed_factor_len; i++ {
		getOpsPermutes(make([]string, i), i-1, 0, &ops_permut, strconv.Itoa(i), true)
	}

	for i := 0; i < len(failed_int_results); i++ {
		factor := strings.Split(failed_factors[i], " ")
		int_factors := []int{}
		for _, i := range factor {
			conv, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			int_factors = append(int_factors, conv)
		}

		permute_count := len(int_factors) - 1
		if failed_int_results[i] == 192 {
			fmt.Println("hey")
		}

		if checkFailedResults(failed_int_results[i], int_factors, ops_permut[strconv.Itoa(permute_count)]) {
			result_sum += failed_int_results[i]
		}
	}

	fmt.Printf("Part two results sum: %v\r\n", result_sum)
}

func checkResults(result int, factors []int, ops [][]string) bool {
	for _, i := range ops {

		end_res := factors[0]

		if len(factors) == 1 {
			return end_res == result
		}

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

func checkFailedResults(result int, factors []int, ops [][]string) bool {
	for _, i := range ops {

		if !slices.Contains(i, "||") {
			continue
		}

		end_res := factors[0]

		if len(factors) == 1 {
			return end_res == result
		}

		for count, j := range i {
			if j == "+" {
				end_res += factors[count+1]
			} else if j == "*" {
				end_res *= factors[count+1]
			} else if j == "||" {
				conv, err := strconv.Atoi(fmt.Sprintf("%v%v", end_res, factors[count+1]))
				if err != nil {
					panic(err)
				}
				end_res = conv
			}
		}

		if end_res == result {
			return true
		}
	}

	return false
}

func getOpsPermutes(data []string, last int, index int, ops_permut *map[string][][]string, loop string, useConcat bool) {
	var ops []string
	if !useConcat {
		ops = []string{"+", "*"}
	} else {
		ops = []string{"+", "*", "||"}
	}
	for i := 0; i < len(ops); i++ {

		data[index] = ops[i]

		if index == last {
			// fmt.Println(data)
			addOutputToMap(ops_permut, loop, strings.Join(data, ":"))
		} else {
			getOpsPermutes(data, last, index+1, ops_permut, loop, useConcat)
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
