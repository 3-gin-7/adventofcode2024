package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	emptyCapacity, totalCapacity := getCapacities(input)
	expanded_string, queue, queueMap := expandMemoryString(input, emptyCapacity, totalCapacity)
	part_two_expanded_string := make([]string, len(expanded_string))

	copy(part_two_expanded_string, expanded_string)

	compact_string := convertToCompactString(expanded_string, &queue)
	checksum := calculateCheckSum(compact_string)
	part_two_string := convertToCompactStringPartTwo(part_two_expanded_string, queueMap)
	part_two_checksum := calculatePartTwoCheckSum(part_two_string)

	fmt.Printf("checksum is: %v\r\n", checksum)
	fmt.Printf("part two checksum is: %v\r\n", part_two_checksum)
}

func calculatePartTwoCheckSum(input []string) uint64 {
	var sum uint64 = 0
	for i := uint64(0); i < uint64(len(input)); i++ {
		str := input[i]
		if str == "." {
			continue
		}

		num, _ := strconv.ParseUint(str, 10, 64)
		sum += i * num
	}

	return sum
}

func convertToCompactStringPartTwo(expanded_string []string, queueMap map[string]Queue) []string {
	// loop over the input array in reverse
	// generate the list of keys
	item := []string{}
	key_list := make([]string, len(queueMap))
	count := 0

	for i := range queueMap {
		key_list[count] = i
		count++
	}

	for i := len(expanded_string) - 1; i > 0; i-- {
		current_char := string(expanded_string[i])

		// if item is empty add current index and continue
		if len(item) == 0 {
			item = append(item, expanded_string[i])
			continue
		}

		// skip the empty spaces
		if current_char == "." && item[0] == "." {
			item = []string{"."}
			continue
		}

		if string(item[0]) == "." && current_char != "." {
			item = []string{current_char}
			continue
		}

		// keep adding the chars to the item until it is not the same char
		if current_char == string(item[0]) {
			item = append(item, current_char)
			continue
		}

		// if here then it is a different non "." char
		item_len := len(item)

		valid_len := getValidKey(&queueMap, item_len)

		// check if map has valid keys
		if valid_len == "" {
			item = []string{current_char}
			continue
		}

		queue := queueMap[valid_len]
		// check if queue has indexes left
		if queue.IsEmpty() {
			item = []string{current_char}
			continue
		}

		empty_index := queue.Pop()

		if empty_index > i {
			item = []string{current_char}
			continue
		}

		queueMap[valid_len] = queue

		// swap the expanded string with item
		curr_index := i
		for j := 0; j < len(item); j++ {
			expanded_string[empty_index+j] = string(item[j])
			expanded_string[curr_index+j+1] = "."
		}

		// calculate the difference in empty spaces left and update the map
		int_valid_len, _ := strconv.Atoi(valid_len)
		empty_diff := strconv.Itoa(int_valid_len - len(item))
		// check if queue is in the map
		if empty_diff != "0" {
			if slices.Contains(key_list, empty_diff) {
				existing_queue := queueMap[empty_diff]
				existing_queue.queue = append(existing_queue.queue, empty_index+len(item))
				sort.Ints(existing_queue.queue)
				existing_queue.curr_index++
				queueMap[empty_diff] = existing_queue
			} else {
				// new queue
				key_list = append(key_list, empty_diff)
				new_queue := Queue{curr_index: 0, fetch_index: 0, queue: []int{empty_index + len(item)}}
				queueMap[empty_diff] = new_queue
			}
		}

		// check if the current index was swapped
		if current_char != expanded_string[i] {
			item = []string{expanded_string[i]}
		} else {
			item = []string{current_char}
		}
	}

	return expanded_string
}

func getValidKey(q *map[string]Queue, item_len int) string {
	min_index := 1000000000000
	chosen_key := ""
	for key, value := range *q {
		if value.IsEmpty() {
			continue
		}
		conv, _ := strconv.Atoi(key)
		if conv >= item_len {
			if min_index > value.queue[0] {
				min_index = value.queue[0]
				chosen_key = key
			}
		}
	}

	return chosen_key
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

func expandMemoryString(input string, empty_capacity int, total_capacity int) ([]string, Queue, map[string]Queue) {
	queueMap := make(map[string]Queue)
	var queue Queue = Queue{curr_index: 0, queue: make([]int, empty_capacity), fetch_index: 0}
	var memory []string = make([]string, total_capacity)

	file_index := 0
	str_index := 0
	for i, str := range strings.Split(input, "") {
		num, _ := strconv.Atoi(str)
		isFileBlock := i%2 == 0

		if !isFileBlock {
			emptyQueue := queueMap[str]
			emptyQueue.queue = append(emptyQueue.queue, str_index)
			emptyQueue.curr_index++

			queueMap[str] = emptyQueue
		}

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

	return memory, queue, queueMap
}

func convertToCompactString(input []string, queue *Queue) []string {

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
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

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
	out := q.queue[0]
	q.queue = q.queue[1:]
	return out
}

func (q *Queue) Add(i int) {
	q.queue[q.curr_index] = i
	q.curr_index++
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
