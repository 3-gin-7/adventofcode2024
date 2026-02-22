package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// byes_to_read := 12
	// max_x := 6
	// max_y := 6
	bytes_to_read := 1024
	max_x := 70
	max_y := 70

	matrix, _, err := readFile(bytes_to_read, max_x, max_y)
	if err != nil {
		panic(err)
	}

	printMatrix(matrix)

	part_one_setps := partOneSteps(matrix, max_x, max_y)

	fmt.Printf("Part one steps: %v\r\n", part_one_setps)

	new_matrix, bytes, err := readFile(-1, max_x, max_y)
	if err != nil {
		panic(err)
	}
	printMatrix(new_matrix)

	part_two_bytes := partTwoBytes(new_matrix, bytes, max_x, max_y)

	fmt.Printf("Part two bytes :%v\r\n", part_two_bytes)
}

func partTwoBytes(new_matrix [][]string, bytes []string, max_x, max_y int) string {
	// since range in 0 to max each row/col count is max + 1 and total elements are the area of the grid
	total_x := max_x + 1
	total_y := max_y + 1
	total := total_x * total_y

	// directions
	directions := []struct{ x, y int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	unions := NewUnion(total)
	// ids of open nodes
	is_open := make([]bool, total)

	// start and end ids
	start_id := getUnionIdFromInts(0, 0, total_x)
	end_id := getUnionIdFromInts(max_x, max_y, total_x)

	// go over the open coords and add mark them
	for y := range new_matrix {
		for x := range new_matrix[y] {
			if new_matrix[y][x] == "." {
				id := getUnionIdFromInts(x, y, total_x)
				is_open[id] = true
			}
		}
	}

	// union connected open nodes
	for y := range new_matrix {
		for x := range new_matrix[y] {
			id := getUnionIdFromInts(x, y, total_x)
			if !is_open[id] {
				continue
			}

			for _, dir := range directions {
				new_x, new_y := x+dir.x, y+dir.y
				// check bounds and if it is open
				if new_x >= 0 && new_x < total_x && new_y >= 0 && new_y < total_y && new_matrix[new_y][new_x] == "." {
					new_id := getUnionIdFromInts(new_x, new_y, total_x)
					// union the nodes
					unions.union(id, new_id)
				}
			}
		}
	}

	for i := 0; i < len(bytes); i++ {
		byte := bytes[len(bytes)-i-1]

		// split the byte into x and y coordinates
		x, y := parseCoords(byte)
		new_matrix[y][x] = "."
		id := getUnionIdFromInts(x, y, total_x)
		is_open[id] = true

		// continue untile start and end are open
		if !is_open[start_id] || !is_open[end_id] {
			continue
		}

		// loop over all dirs and check if any are in nodes
		for _, dir := range directions {
			new_x, new_y := x+dir.x, y+dir.y
			// check bounds
			if new_x < 0 || new_x > max_x || new_y < 0 || new_y > max_y {
				continue
			}

			// check if wall
			if new_matrix[new_y][new_x] == "#" {
				continue
			}

			// union id of new node
			new_id := getUnionIdFromInts(new_x, new_y, total_x)

			// union the nodes
			unions.union(id, new_id)
		}

		// check if start and end are connected
		if unions.find(start_id) == unions.find(end_id) {
			return byte
		}
	}

	return ""
}

// bfs search for shortest path
func partOneSteps(matrix [][]string, max_x, max_y int) int {
	start_x := 0
	start_y := 0

	end_x := max_x
	end_y := max_y

	queue := make(StepQueue, 0)
	queue.Push(Queue{coord: fmt.Sprintf("%v,%v", start_x, start_y), steps: 0})
	matrix[start_y][start_x] = "#"

	for len(queue) > 0 {
		current := queue.Pop()
		x, y := parseCoords(current.coord)

		// check if at the end
		if x == end_x && y == end_y {
			fmt.Printf("queue len: %v\r\n", len(queue))
			return current.steps
		}

		directions := []struct{ x, y int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		for _, dir := range directions {
			new_x, new_y := x+dir.x, y+dir.y
			// check bounds
			if new_x >= 0 && new_x <= max_x && new_y >= 0 && new_y <= max_y && matrix[new_y][new_x] == "." {
				queue.Push(Queue{coord: fmt.Sprintf("%v,%v", new_x, new_y), steps: current.steps + 1})
				matrix[new_y][new_x] = "O"
			}
		}
	}

	fmt.Println(start_x, start_y, end_x, end_y)
	return 0
}

func parseCoords(coord string) (int, int) {
	parts := strings.Split(coord, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

func getUnionIdFromStr(coords string, max int) int {
	x, y := parseCoords(coords)
	return getUnionIdFromInts(x, y, max)
}

func getUnionIdFromInts(x, y, max int) int {
	return y*max + x
}

func readFile(bytes_to_read, max_x, max_y int) ([][]string, []string, error) {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	type point struct{ x, y int }
	points := make([]point, 0)
	bytes_read := 0
	var bytes []string

	for sc.Scan() {
		if bytes_to_read >= 0 && bytes_read >= bytes_to_read {
			break
		}

		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		// splint and convert
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid coords: %q", line)
		}

		x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("x not a num %q: %w", parts[0], err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, nil, fmt.Errorf("y not a num %q: %w", parts[1], err)
		}
		// check bounds
		if x < 0 || x > max_x || y < 0 || y > max_y {
			return nil, nil, fmt.Errorf("(%d,%d) out of bounds for max (%d,%d)", x, y, max_x, max_y)
		}

		points = append(points, point{x: x, y: y})
		bytes = append(bytes, line)
		bytes_read++
	}

	if err := sc.Err(); err != nil {
		return nil, nil, err
	}

	rows := max_y + 1
	cols := max_x + 1
	matrix := make([][]string, rows)
	for r := range matrix {
		matrix[r] = make([]string, cols)
		for c := range matrix[r] {
			matrix[r][c] = "."
		}
	}

	for _, p := range points {
		matrix[p.y][p.x] = "#"
	}

	return matrix, bytes, nil
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}

type Queue struct {
	coord string
	steps int
}

type StepQueue []Queue

func (q *StepQueue) Push(val Queue) {
	*q = append(*q, val)
}

func (q *StepQueue) Pop() Queue {
	if len(*q) == 0 {
		return Queue{}
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *StepQueue) IsEmpty() bool {
	return len(*q) == 0
}

type Union struct {
	parent []int
}

func NewUnion(len int) *Union {
	union := &Union{
		parent: make([]int, len),
	}

	// set parent to itself
	for i := 0; i < len; i++ {
		union.parent[i] = i
	}

	return union
}

// get the id of root
func (union *Union) find(i int) int {
	if union.parent[i] != i {
		union.parent[i] = union.find(union.parent[i])
	}
	return union.parent[i]
}

// union two nodes
func (union *Union) union(a, b int) {
	root_a := union.find(a)
	root_b := union.find(b)
	if root_a == root_b {
		return
	}

	union.parent[root_a] = root_b
}
