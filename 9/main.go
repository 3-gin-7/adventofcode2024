package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	emptyCapacity, totalCapacity := getCapacities(input)
	expanded_string, queue := expandMemoryString(input, emptyCapacity, totalCapacity)

	compact_string := convertToCompactString(expanded_string, &queue)
	checksum := calculateCheckSum(compact_string)

	fmt.Printf("checksum is: %v\r\n", checksum)
}

func getCapacities(input string) (int, int) {
	empty_capacity := 0
	total_capacity := 0
	for i, val := range strings.Split(input, "") {
		num, _ := strconv.Atoi(val)
		total_capacity += num
		if i%2 == 1 {
			empty_capacity += num
		}
	}
	return empty_capacity, total_capacity
}

func calculateCheckSum(input []string) int {
	sum := 0
	for i, str := range input {
		if str == "." {
			break
		}

		num, _ := strconv.Atoi(str)
		sum += i * num
	}

	return sum
}

func expandMemoryString(input string, empty_capacity int, total_capacity int) ([]string, Queue) {
	var queue Queue = Queue{curr_index: 0, queue: make([]int, empty_capacity), fetch_index: 0}
	var memory []string = make([]string, total_capacity)

	file_index := 0
	str_index := 0
	for i, str := range strings.Split(input, "") {
		num, _ := strconv.Atoi(str)
		isFileBlock := i%2 == 0

		// run the number of times
		for j := 0; j < num; j++ {
			// file block
			if isFileBlock {
				memory[str_index] = strconv.Itoa(file_index)
			} else {
				memory[str_index] = "."
				queue.Add(str_index)
			}

			str_index++
		}

		if isFileBlock {
			file_index++
		}
	}

	return memory, queue
}

func convertToCompactString(input []string, queue *Queue) []string {
	// var sb strings.Builder

	for i := len(input) - 1; i >= 0; i-- {
		if (*queue).fetch_index > len((*queue).queue)-1 {
			break
		}

		if string(input[i]) != "." {
			empty_index := (*queue).Pop()
			if empty_index == 0 {
				break
			}

			if i < empty_index {
				break
			}

			input[empty_index] = input[i]
			input[i] = "."
		}
	}

	return input
}

func readFile() string {
	FILE_NAME := "test.txt"
	// FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)

	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(fi)
	output := ""
	for sc.Scan() {
		output = sc.Text()
	}

	return output
}

type Queue struct {
	queue       []int
	curr_index  int
	fetch_index int
}

func (q *Queue) Pop() int {
	out := q.queue[q.fetch_index]
	q.fetch_index++
	return out
}

func (q *Queue) Add(i int) {
	q.queue[q.curr_index] = i
	q.curr_index++
}
