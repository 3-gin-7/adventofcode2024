package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	matrix, trail_starts := readFile()

	// fmt.Println(matrix)
	// fmt.Println(trail_starts)

	// loop over the trail starts
	// checkIfFullTrail
	// get the starting point "0" coords
	output := 0
	trail_count := 0
	trailCount := make(map[string][]string)
	for _, i := range trail_starts {
		trailCount[i] = []string{}
		val := getTrailCount(i, matrix, []string{}, &trailCount)
		output += val
	}

	for _, value := range trailCount {
		trail_count += len(value)
	}

	fmt.Println(trail_count)
	// fmt.Println(output)
}

func getTrailCount(i string, matrix [][]string, history []string, uniqueTrails *map[string][]string) int {
	history = append(history, i)
	trail_count := 0

	// check up coords
	up_coords := getDirectionCoords(i, "up", matrix)
	if up_coords == "x" {
		next_coords := getNextCoords(i, "up")
		trail_count++
		updateTrailCount(history, uniqueTrails, next_coords)
	} else if up_coords != "" {
		trail_count += getTrailCount(up_coords, matrix, history, uniqueTrails)
	}

	// check right coords
	right_coords := getDirectionCoords(i, "right", matrix)
	if right_coords == "x" {
		next_coords := getNextCoords(i, "right")
		updateTrailCount(history, uniqueTrails, next_coords)
		trail_count++
	} else if right_coords != "" {
		trail_count += getTrailCount(right_coords, matrix, history, uniqueTrails)
	}

	// check down coords
	down_coords := getDirectionCoords(i, "down", matrix)
	if down_coords == "x" {
		next_coords := getNextCoords(i, "down")
		updateTrailCount(history, uniqueTrails, next_coords)
		trail_count++
	} else if down_coords != "" {
		trail_count += getTrailCount(down_coords, matrix, history, uniqueTrails)
	}

	// check left coords
	left_coords := getDirectionCoords(i, "left", matrix)
	if left_coords == "x" {
		next_coords := getNextCoords(i, "left")
		updateTrailCount(history, uniqueTrails, next_coords)
		trail_count++
	} else if left_coords != "" {
		trail_count += getTrailCount(left_coords, matrix, history, uniqueTrails)
	}

	return trail_count
}

func updateTrailCount(history []string, uniqueTrails *map[string][]string, next_coords string) {
	if len(history) == 9 {
		trail_start := history[0]
		trails := (*uniqueTrails)[trail_start]
		if !slices.Contains(trails, next_coords) {
			trails = append(trails, next_coords)
			(*uniqueTrails)[trail_start] = trails
		}
	}
}

func getNextCoords(i, s string) string {
	coords := strings.Split(i, ",")
	str1 := coords[0]
	str2 := coords[1]

	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)

	if s == "up" {
		num1--
	} else if s == "right" {
		num2++
	} else if s == "down" {
		num1++
	} else if s == "left" {
		num2--
	}

	return fmt.Sprintf("%v,%v", num1, num2)
}

func getDirectionCoords(i, direction string, matrix [][]string) string {
	coords := strings.Split(i, ",")
	str1 := coords[0]
	str2 := coords[1]

	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)
	current_val, _ := strconv.Atoi(matrix[num1][num2])

	if direction == "up" {
		num1--
		if num1 < 0 {
			return ""
		}
	} else if direction == "right" {
		num2++
		if num2 >= len(matrix[0]) {
			return ""
		}
	} else if direction == "down" {
		num1++
		if num1 >= len(matrix[0]) {
			return ""
		}
	} else if direction == "left" {
		num2--
		if num2 < 0 {
			return ""
		}
	}

	// check if the next value is increasing
	next_val, _ := strconv.Atoi(matrix[num1][num2])
	if next_val-current_val == 1 {
		if next_val == 9 {
			return "x"
		}

		return fmt.Sprintf("%v,%v", num1, num2)
	}

	return ""
}

func readFile() ([][]string, []string) {
	matrix := [][]string{}
	trailheads := []string{}

	// FILE_NAME := "test.txt"
	// FILE_NAME := "tmp.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)

	if err != nil {
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)

	outter_count := 0
	for sc.Scan() {
		line := sc.Text()
		matrix = append(matrix, strings.Split(line, ""))

		for idx, i := range strings.Split(line, "") {
			if i == "0" {
				trailheads = append(trailheads, fmt.Sprintf("%v,%v", outter_count, idx))
			}
		}
		outter_count++
	}

	return matrix, trailheads
}
