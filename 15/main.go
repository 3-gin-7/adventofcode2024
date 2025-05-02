package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boxes, walls, part_two_boxes, part_two_walls, moves, pos := readFile()

	fmt.Printf("part one output: %v \r\n", getPartOneCount(boxes, walls, moves, pos))
	fmt.Printf("part two output: %v \r\n", getPartTwoCount(part_two_boxes, part_two_walls, moves, pos))
}

func getPartTwoCount(boxes, walls map[string]string, moves []string, pos string) int {
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

		// print out the move stuff
		// printGrid(boxes, walls, pos)
		// fmt.Printf("Move #%v: %v \r\n", i, m)

		x, y := getIntPos(pos)
		tmp_pos := fmt.Sprintf("%v,%v", x+to_row, y+to_col)

		box, ok_box := boxes[tmp_pos]
		start_col, end_col := 0, 0

		if ok_box {
			if box == "]" {
				start_col = y - 1
				end_col = y
			} else if box == "[" {
				end_col = y + 1
				start_col = y
			}
		}

		should_move, m_boxes := move(to_row, to_col, tmp_pos, "@", start_col, end_col, walls, boxes, []string{})

		if should_move {
			moved_cache := make(map[string]string)
			for i := len(m_boxes) - 1; i >= 0; i-- {
				key1 := m_boxes[i]
				val := boxes[key1]

				_, ok_cache := moved_cache[key1]
				if ok_cache {
					continue
				}

				key2 := ""
				new_key := ""
				_x, _y := getIntPos(key1)

				if m == "<" || m == ">" {
					key2 = fmt.Sprintf("%v,%v", _x, _y+to_col)
					new_key = fmt.Sprintf("%v,%v", _x, _y+(to_col*2))

					a := boxes[key1]
					b := boxes[key2]
					delete(boxes, key1)
					delete(boxes, key2)

					moved_cache[key1] = ""
					moved_cache[key2] = ""

					boxes[key2] = a
					boxes[new_key] = b

				} else {
					new_key2 := ""
					if val == "]" {
						key2 = fmt.Sprintf("%v,%v", _x, _y-1)
						new_key2 = fmt.Sprintf("%v,%v", _x+to_row, _y-1)
					} else {
						key2 = fmt.Sprintf("%v,%v", _x, _y+1)
						new_key2 = fmt.Sprintf("%v,%v", _x+to_row, _y+1)
					}

					new_key1 := fmt.Sprintf("%v,%v", _x+to_row, _y)

					a := boxes[key1]
					b := boxes[key2]

					moved_cache[key1] = ""
					moved_cache[key2] = ""

					delete(boxes, key1)
					delete(boxes, key2)

					// if a == "" || b == "" {
					// 	fmt.Println("yo")
					// }

					boxes[new_key1] = a
					boxes[new_key2] = b
				}

			}
			pos = tmp_pos
		}

	}

	part_two_output := 0

	for key, val := range boxes {
		if val == "[" {
			row, col := getIntPos(key)
			part_two_output += (row * 100) + col
		}
	}

	return part_two_output
}

func printGrid(boxes, walls map[string]string, pos string) {
	for i := 0; i < 50; i++ {
		line := ""
		for j := 0; j < 100; j++ {
			if i == 0 || i == 100 {
				line = "####################################################################################################"
				break
			}

			key := fmt.Sprintf("%v,%v", i, j)
			box, ok_box := boxes[key]
			_, ok_wall := walls[key]

			if ok_box {
				line += box
			} else if ok_wall {
				line += "#"
			} else if key == pos {
				line += "@"
			} else {
				line += "."
			}

		}

		fmt.Println(line)
	}
	fmt.Println()
}

func move(row, col int, pos, prev_item string, start_col, end_col int, walls, boxes map[string]string, moved_boxes []string) (bool, []string) {
	should_move := false

	// moving left/right
	if row == 0 {
		_, ok_wall := walls[pos]
		if ok_wall {
			return false, moved_boxes
		}

		box, ok_box := boxes[pos]

		if !ok_box {
			// has to be an open space
			return true, moved_boxes
		}

		moved_boxes = append(moved_boxes, pos)

		if box == "]" {

			x, y := getIntPos(pos)
			pos = fmt.Sprintf("%v,%v", x, y-2)
			should_move, moved_boxes = move(row, col, pos, box, start_col, end_col, walls, boxes, moved_boxes)
		} else if box == "[" {
			x, y := getIntPos(pos)
			pos = fmt.Sprintf("%v,%v", x, y+2)
			should_move, moved_boxes = move(row, col, pos, box, start_col, end_col, walls, boxes, moved_boxes)
		} else {
			panic("Invalid box")
		}

		return should_move, moved_boxes
	} else if col == 0 {
		// moving up/down
		box, ok_box := boxes[pos]
		_, ok_wall := walls[pos]

		if ok_wall {
			return false, moved_boxes
		}

		// check if the empty space
		if !ok_box {
			return true, moved_boxes
		}

		// has to be a box
		moved_boxes = append(moved_boxes, pos)

		move_flag := true
		x, y := getIntPos(pos)
		x += row

		if prev_item != "@" {
			if box == "[" {
				start_col = y
				end_col = y + 1
			} else if box == "]" {
				end_col = y
				start_col = y - 1
			}
		}

		for i := start_col; i <= end_col; i++ {
			// inc pos
			pos = fmt.Sprintf("%v,%v", x, i)
			move_flag, moved_boxes = move(row, col, pos, box, start_col, end_col, walls, boxes, moved_boxes)

			if !move_flag {
				return false, moved_boxes
			}
		}

		return true, moved_boxes

	} else {
		panic("invalid move")
	}
}

func getPartOneCount(boxes, walls map[string]string, moves []string, pos string) int {
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
	for key := range boxes {
		// 100 * row + col
		row, col := getIntPos(key)
		part_one_output += (row * 100) + col
	}

	return part_one_output
}

func getIntPos(key string) (int, int) {

	split := strings.Split(key, ",")
	row, _ := strconv.Atoi(split[0])
	col, _ := strconv.Atoi(split[1])

	return row, col
}

func readFile() (map[string]string, map[string]string, map[string]string, map[string]string, []string, string) {
	boxes := make(map[string]string)
	walls := make(map[string]string)
	part_two_boxes := make(map[string]string)
	part_two_walls := make(map[string]string)

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

		// part two input
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")

		line_split = strings.Split(line, "")
		for col, l := range line_split {
			key := fmt.Sprintf("%v,%v", row, col)

			if l == "[" || l == "]" {
				part_two_boxes[key] = l
			}
			if l == "@" {
				pos = key
			}

			if l == "#" {
				part_two_walls[key] = key
			}
		}

		row++
	}

	return boxes, walls, part_two_boxes, part_two_walls, moves, pos
}
