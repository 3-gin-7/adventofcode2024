package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boxes, walls, moves, pos := readFile()

	for _, m := range moves {
		to_row := 0
		to_col := 0
		if m == "<" {
			// move left
			to_col = -1
		} else if m == "^" {
			// move up
			to_row = -1
		} else if m == "v" {
			// move down
			to_row = 1
		} else if m == ">" {
			// move right
			to_col = 1
		} else {
			panic("wrong move")
		}

		should_move := false
		moved_boxes := []string{}
		new_pos := pos

		for {
			split := strings.Split(new_pos, ",")
			row, _ := strconv.Atoi(split[0])
			col, _ := strconv.Atoi(split[1])

			row += to_row
			col += to_col

			new_pos = fmt.Sprintf("%v,%v", row, col)

			key, ok_box := boxes[new_pos]
			_, ok_wall := walls[new_pos]

			if ok_box {
				moved_boxes = append(moved_boxes, key)
				continue
			} else if ok_wall {
				// border
				break
			} else {
				// open space
				should_move = true
				break
			}
		}

		if should_move {
			split := strings.Split(pos, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			should_move = false

			// update boxes position
			for i, m := range moved_boxes {
				x_box, y_box := getIntPos(m)
				new_key := fmt.Sprintf("%v,%v", x_box+to_row, y_box+to_col)
				og_key := fmt.Sprintf("%v,%v", x_box, y_box)

				// only delete the first moved box
				if i == 0 {
					delete(boxes, og_key)
				}

				_, ok := boxes[new_key]
				if !ok {
					boxes[new_key] = new_key
				}
			}

			// update robot pos
			pos = fmt.Sprintf("%v,%v", x+to_row, y+to_col)
		}
	}

	// fmt.Println(boxes)
	part_one_output := 0
	for key, _ := range boxes {
		// 100 * row + col
		row, col := getIntPos(key)
		part_one_output += (row * 100) + col
	}

	fmt.Printf("part one output: %v \r\n", part_one_output)
}

func getIntPos(key string) (int, int) {

	split := strings.Split(key, ",")
	row, _ := strconv.Atoi(split[0])
	col, _ := strconv.Atoi(split[1])

	return row, col
}

func readFile() (map[string]string, map[string]string, []string, string) {
	boxes := make(map[string]string)
	walls := make(map[string]string)

	pos := ""
	moves := []string{}

	// FILE_NAME := "test.txt"
	// FILE_NAME := "tmp.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()

	moves_flag := false
	sc := bufio.NewScanner(fi)
	row := 0
	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			moves_flag = true
			continue
		}

		line_split := strings.Split(line, "")

		for col, l := range line_split {
			key := fmt.Sprintf("%v,%v", row, col)

			if moves_flag {
				moves = append(moves, l)
			}
			if l == "O" {
				boxes[key] = key
			}
			if l == "@" {
				pos = key
			}

			if l == "#" {
				walls[key] = key
			}
		}

		row++
	}

	return boxes, walls, moves, pos
}
