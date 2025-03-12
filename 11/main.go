package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this is super slow
// need to think of another way of doing this
func main() {
	number_of_blinks := 25

	stones := readFile()
	fmt.Println(stones.listToSlice())

	for i := 0; i < number_of_blinks; i++ {
		for j := 0; j < stones.length; j++ {

			node := stones.getAtIndex(j)
			transformed_stones := applyRules(node.data)
			if len(transformed_stones) == 1 {
				node.data = strconv.Itoa(transformed_stones[0])
			} else {
				stones.insertAtIndex([]string{strconv.Itoa(transformed_stones[0]), strconv.Itoa(transformed_stones[1])}, j)
				j++
			}
		}
	}
	fmt.Println(stones.length)
}

func applyRules(data string) []int {
	int_data, _ := strconv.Atoi(data)
	if int_data == 0 {
		return []int{1}
	}

	node_length := len(data)
	if node_length%2 == 0 {
		first_part, _ := strconv.Atoi(data[:node_length/2])
		second_part, _ := strconv.Atoi(data[node_length/2:])

		return []int{first_part, second_part}
	}

	return []int{int_data * 2024}
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
