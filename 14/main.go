package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	robots := readFile()

	moves := 100
	x_max := 101
	y_max := 103
	// x_max := 11
	// y_max := 7

	quadrant_1 := 0
	quadrant_2 := 0
	quadrant_3 := 0
	quadrant_4 := 0

	output := [][]int{}
	for i := 0; i < y_max; i++ {
		output = append(output, []int{})
		for j := 0; j < x_max; j++ {
			output[i] = append(output[i], 0)
		}
	}

	// for _, o := range output {
	// 	fmt.Println(o)
	// }

	for _, robot := range robots {
		robot.calculatePos(moves, x_max, y_max)
		robot_pos := robot.finalPos

		// t := x_max / 2
		// fmt.Println(t)

		// hate this
		x := robot_pos[0]
		y := robot_pos[1]
		output[y][x]++

		if x == x_max/2 || y == y_max/2 {
			continue
		}

		if x < x_max/2 {
			// either 1 or 3
			if y < y_max/2 {
				quadrant_1++
			} else {
				quadrant_3++
			}
		} else {
			// either 2 or 4
			if y < y_max/2 {
				quadrant_2++
			} else {
				quadrant_4++
			}
		}
	}

	fmt.Println("output")
	// for _, o := range output {
	// 	fmt.Println(o)
	// }

	fmt.Printf("part 1: %d\n", quadrant_1*quadrant_2*quadrant_3*quadrant_4)

	// fmt.Println(quadrant_1, quadrant_2, quadrant_3, quadrant_4)
}

func readFile() []*Robot {
	output := []*Robot{}
	FILE_NAME := "input.txt"
	// FILE_NAME := "test.txt"
	// FILE_NAME := "tmp.txt"

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		robot := &Robot{}
		line := scanner.Text()
		split := strings.Split(line, " ")
		pos := strings.Split(split[0], "=")
		vel := strings.Split(split[1], "=")
		posInt := strings.Split(pos[1], ",")
		velInt := strings.Split(vel[1], ",")

		robot.startPos = []int{stringToInt(posInt[0]), stringToInt(posInt[1])}
		robot.velocity = []int{stringToInt(velInt[0]), stringToInt(velInt[1])}
		output = append(output, robot)
	}

	return output
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return num
}

type Robot struct {
	startPos []int
	velocity []int
	finalPos []int
}

func (r *Robot) calculatePos(move int, x_max int, y_max int) {
	offset_x := r.startPos[0]
	offset_y := r.startPos[1]

	velocity_x := r.velocity[0]
	velocity_y := r.velocity[1]

	final_x := 0
	final_y := 0

	// calculate x
	if velocity_x > 0 {
		final_x = (move*velocity_x + offset_x) % x_max
	} else {
		total_move := math.Abs(float64(move * velocity_x))
		final_x = (offset_x - int(total_move)) % x_max
		if final_x < 0 {
			final_x = x_max + final_x
		}
	}

	// calculate y
	if velocity_y > 0 {
		final_y = (move*velocity_y + offset_y) % y_max
	} else {
		total_move := math.Abs(float64(move * velocity_y))
		final_y = (offset_y - int(total_move)) % y_max
		if final_y < 0 {
			final_y = y_max + final_y
		}
	}

	r.finalPos = []int{final_x, final_y}
}
