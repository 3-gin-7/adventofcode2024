package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := readFile()
	fmt.Println(stones.listToSlice())
	seen_numbers := make(map[string][]string)

	for i := 0; i < 75; i++ {
		fmt.Printf("iteration %v: stones count: %v\r\n", i, stones.length)
		for j := 0; j < stones.length; j++ {
			node := stones.getAtIndex(j)
			transformed_stones := strings.Split(node.data, " ")

			value, ok := seen_numbers[node.data]
			if ok {
				transformed_stones = value
			} else {
				transformed_stones = applyRulesTwo(transformed_stones)
				seen_numbers[node.data] = transformed_stones
			}

			stones.insertRangeAtIndex(transformed_stones, j)
			j += len(transformed_stones) - 1
		}
	}

	fmt.Println(stones.length)
}

func test(blink int, stones string, stone_count int) int {
	if blink == 25 {
		return stone_count
	}

	transformed_stone := applyRules(stones)
	if len(transformed_stone) > 1 {
		stone_count++
	}

	for _, i := range transformed_stone {
		if i != "1" && i != "0" {
			stone_count = test(blink+1, i, stone_count)
		}
	}

	return stone_count
}

func applyRulesTwo(data []string) []string {
	output := []string{}
	for _, i := range data {
		new_data := applyRules(i)
		output = append(output, new_data...)
	}

	return output
}

func applyRules(data string) []string {
	int_data, _ := strconv.Atoi(data)
	if int_data == 0 {
		return []string{"1"}
	}

	node_length := len(data)
	if node_length%2 == 0 {
		first_part, _ := strconv.Atoi(data[:node_length/2])
		second_part, _ := strconv.Atoi(data[node_length/2:])

		return []string{strconv.Itoa(first_part), strconv.Itoa(second_part)}
	}

	return []string{strconv.Itoa(int_data * 2024)}
}

func readFile() *LinkedList {
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"
	fi, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)
	sc.Scan()

	list := &LinkedList{}

	for _, stone := range strings.Split(sc.Text(), " ") {
		list.insert(stone)
	}

	return list
}

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) insert(data string) {
	list.length++
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (list *LinkedList) insertAtIndex(data []string, index int) {
	if list.length <= index {
		panic("index is out of range")
	}
	list.length++

	insert_node := list.head
	for i := 0; i < index; i++ {
		insert_node = insert_node.next
	}

	future_next := insert_node.next
	insert_node.data = data[0]
	insert_node.next = &Node{data: data[1], next: future_next}
}

func (list *LinkedList) insertRangeAtIndex(data []string, index int) {
	if list.length <= index {
		panic("index is out of range")
	}
	list.length += len(data) - 1

	insert_node := list.head
	for i := 0; i < index; i++ {
		insert_node = insert_node.next
	}

	insert_node.data = data[0]
	link_node := insert_node.next
	for i := 1; i < len(data); i++ {
		new_node := &Node{data: data[i], next: link_node}
		insert_node.next = new_node
		insert_node = new_node
	}

	insert_node.next = link_node
}

func (list *LinkedList) listToSlice() []string {
	output := []string{}

	current := list.head

	for i := 0; i < list.length; i++ {
		output = append(output, current.data)
		current = current.next
	}

	return output
}

func (list *LinkedList) getAtIndex(index int) *Node {
	if index >= list.length {
		panic("index out of the bounds")
	}

	node := list.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (list *LinkedList) toList() []string {
	output := make([]string, list.length)
	for i := 0; i < list.length; i++ {
		node := list.getAtIndex(i)
		output[i] = node.data
	}

	return output
}
