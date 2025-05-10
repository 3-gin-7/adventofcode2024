package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	start, end, nodes := readFile()

	part_one_cost := getPartOnePaths(nodes, start, end)

	fmt.Printf("Part one count: %v\r\n", part_one_cost)
}

func getPartOnePaths(graph map[string]string, start, end string) int {
	dists := make(map[string]int)
	for key, _ := range graph {
		dists[key] = math.MaxInt64
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Edge{Cost: 0, Coord: start, Direction: ">"})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Edge)
		dist := item.Cost
		coords := item.Coord
		direction := item.Direction
		if coords == end {
			return dist
		}

		x, y := getIntCoords(coords)

		up_coords := getStrCoords(x-1, y)
		down_coords := getStrCoords(x+1, y)
		left_coords := getStrCoords(x, y-1)
		right_coords := getStrCoords(x, y+1)

		_, ok_up := graph[up_coords]
		_, ok_down := graph[down_coords]
		_, ok_left := graph[left_coords]
		_, ok_right := graph[right_coords]

		if ok_up {
			cost := getDirectionIncrease(direction, "^")
			if dist+cost < dists[up_coords] {
				dists[up_coords] = dist + cost
				heap.Push(pq, &Edge{Cost: dists[up_coords], Direction: "^", Coord: up_coords})
			}
		}
		if ok_down {
			cost := getDirectionIncrease(direction, "v")
			if dist+cost < dists[down_coords] {
				dists[down_coords] = dist + cost
				heap.Push(pq, &Edge{Cost: dists[down_coords], Direction: "v", Coord: down_coords})
			}
		}
		if ok_left {
			cost := getDirectionIncrease(direction, "<")
			if dist+cost < dists[left_coords] {
				dists[left_coords] = dist + cost
				heap.Push(pq, &Edge{Cost: dists[left_coords], Direction: "<", Coord: left_coords})
			}
		}
		if ok_right {
			cost := getDirectionIncrease(direction, ">")
			if dist+cost < dists[right_coords] {
				dists[right_coords] = dist + cost
				heap.Push(pq, &Edge{Cost: dists[right_coords], Direction: ">", Coord: right_coords})
			}
		}

	}

	return 666
}

func readFile() (string, string, map[string]string) {
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

	nodes := make(map[string]string)

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)

	endpoint := ""
	startpoint := ""

	count := 0
	for sc.Scan() {
		line := sc.Text()
		for i, j := range strings.Split(line, "") {
			coords := getStrCoords(count, i)
			if j == "." {
				nodes[coords] = "."
			} else if j == "S" {
				nodes[coords] = "."
				startpoint = getStrCoords(count, i)
			} else if j == "E" {
				nodes[coords] = "."
				endpoint = getStrCoords(count, i)
			}
		}
		count++
	}

	return startpoint, endpoint, nodes
}

func getStrCoords(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func getIntCoords(pos string) (int, int) {
	split := strings.Split(pos, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	return x, y
}

func getDirectionIncrease(curr_direction, next_direction string) int {
	if next_direction == curr_direction {
		return 1
	}

	if curr_direction == "^" || curr_direction == "v" {
		if next_direction == "<" || next_direction == ">" {
			return 1001
		} else {
			return 2001
		}
	} else if curr_direction == "<" || curr_direction == ">" {
		if next_direction == "^" || next_direction == "v" {
			return 1001
		} else {
			return 2001
		}
	} else {
		panic("Invalid direction")
	}
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Edge)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

type Edge struct {
	Index     int
	Coord     string
	Direction string
	Cost      int
}
