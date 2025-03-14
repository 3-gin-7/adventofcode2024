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
	count := 0

	for j := 0; j < stones.length; j++ {
		count += test(0, stones.getAtIndex(j).data, 1)
	}
	fmt.Println(count)
}

func test(blink int, stones string, stone_count int) int {
	if blink == 25 {
		return stone_count
	}

	transformed_stone := applyRulesTwo(stones)
	if len(transformed_stone) > 1 {
		stone_count++
	}

	for _, i := range transformed_stone {
		stone_count = test(blink+1, i, stone_count)
	}

	return stone_count
}

func applyRulesTwo(data string) []string {
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
	FILE_NAME := "test.txt"
	// FILE_NAME := "input.txt"
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
