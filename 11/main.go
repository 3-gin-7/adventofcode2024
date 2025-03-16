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
	max_depth := 75
	seen_numbers := make(map[string]int)
	stone_count := 0

	list_count := stones.length

	for j := 0; j < list_count; j++ {
		fmt.Printf("processing %v out of %v\r\n", j, list_count)
		// get the node
		node := stones.getAtIndex(j)
		stone_count += generateBranches(max_depth+1, 0, node, stones, seen_numbers)
	}

	fmt.Println(stone_count)
}

func generateBranches(max_depth, current_depth int, node *Node, stones *LinkedList, seen_numbers map[string]int) int {
	stone_count := 0
	current_depth++
	node.stone_count = node.getChildrenStoneCount()

	key := fmt.Sprintf("%v,%v", node.data, current_depth)

	if current_depth == max_depth {
		stone_count++
		return stone_count
	}

	val, ok := seen_numbers[key]

	if ok {
		stone_count += val
		return stone_count
	}

	trans_stones := applyRules(node.data)
	node.left = &Node{data: trans_stones[0], parent: node, stone_count: 1}

	if len(trans_stones) > 1 {
		node.right = &Node{data: trans_stones[1], parent: node, stone_count: 1}
	}

	for i := 0; i < len(trans_stones); i++ {
		if i == 0 {
			stone_count += generateBranches(max_depth, current_depth, node.left, stones, seen_numbers)
		} else {
			stone_count += generateBranches(max_depth, current_depth, node.right, stones, seen_numbers)
		}
	}

	seen_numbers[key] = stone_count
	return stone_count
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
	data        string
	parent      *Node // parent node
	next        *Node // next in the list
	left        *Node // left child
	right       *Node // right child
	stone_count int
	depth       int
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

func (list *LinkedList) listToSlice() []string {
	output := []string{}

	current := list.head

	for i := 0; i < list.length; i++ {
		output = append(output, current.data)
		current = current.next
	}

	return output
}

func (node *Node) getChildrenStoneCount() int {
	count := 0
	if node.left != nil {
		count += node.left.stone_count
	}

	if node.right != nil {
		count += node.right.stone_count
	}

	return count
}

func (node *Node) getStoneCountForDepth(depth int) int {

	return recursiveDepthStoneCount(depth, 0, node)
}

func recursiveDepthStoneCount(depth, count int, node *Node) int {
	if depth == 0 {
		return 1
	}

	depth--

	if node.left != nil {
		count += recursiveDepthStoneCount(depth, count, node.left)
	}

	if node.right != nil {
		count += recursiveDepthStoneCount(depth, count, node.left)
	}
	return count
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
