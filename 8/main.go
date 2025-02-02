package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	antennas := make(map[string]int)
	antennas_part_two := make(map[string]int)
	coords, x_max, y_max := readFile()

	// loop over the corrds
	for _, coord_list := range coords {
		// check the coord list length
		if len(coord_list) < 2 {
			continue
		}

		// match every antenna of the same frequency (excluding itself)
		for x := 0; x < len(coord_list); x++ {
			for y := x + 1; y < len(coord_list); y++ {
				// get the increasing and decreasing difference
				first_diff := getDiff(coord_list[x], coord_list[y])
				second_diff := getDiff(coord_list[y], coord_list[x])

				// get the coords from the difference
				inc_coords, err1 := getCoords(coord_list[x], first_diff, x_max, y_max)
				desc_coords, err2 := getCoords(coord_list[y], second_diff, x_max, y_max)

				if err1 == nil {
					antennas[inc_coords]++
					antennas_part_two[inc_coords]++
				}

				if err2 == nil {
					antennas[desc_coords]++
					antennas_part_two[desc_coords]++
				}

				// part two - loop adding/subtracting the difference until out of bounds
				rec_inc_coords := inc_coords
				if inc_coords != "" {
					for {
						out, err := getCoords(rec_inc_coords, first_diff, x_max, y_max)
						if err != nil {
							break
						}
						antennas_part_two[out]++
						rec_inc_coords = out
					}
				}

				rec_desc_coords := desc_coords
				if desc_coords != "" {
					for {
						out, err := getCoords(rec_desc_coords, second_diff, x_max, y_max)
						if err != nil {
							break
						}
						antennas_part_two[out]++
						rec_desc_coords = out
					}
				}
			}
		}
	}

	fmt.Printf("part one is: %v\r\n", len(antennas))

	// adding the existing antennas to the count
	for _, arr := range coords {
		for _, i := range arr {
			antennas_part_two[i]++
		}
	}

	fmt.Printf("part two len is:%v\r\n", len(antennas_part_two))
}

// calculate the difference/gradient between coords
func getDiff(s1, s2 string) []string {
	// split the numbers
	one := strings.Split(s1, ",")
	two := strings.Split(s2, ",")
	t_x1, t_y1 := one[0], one[1]
	t_x2, t_y2 := two[0], two[1]

	// convert str to int
	x1, _ := strconv.Atoi(t_x1)
	y1, _ := strconv.Atoi(t_y1)
	x2, _ := strconv.Atoi(t_x2)
	y2, _ := strconv.Atoi(t_y2)

	x_diff := x1 - x2
	y_diff := y1 - y2

	return []string{strconv.Itoa(x_diff), strconv.Itoa(y_diff)}
}

// add gradient to the coords
func getCoords(og_coords string, diff []string, x_max, y_max int) (string, error) {
	split := strings.Split(og_coords, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	x_diff, _ := strconv.Atoi(diff[0])
	y_diff, _ := strconv.Atoi(diff[1])

	x = x + x_diff
	y = y + y_diff

	if x < 0 || y < 0 {
		return "", errors.New("invalid coords")
	}

	if x >= x_max || y >= y_max {
		return "", errors.New("invalid coords")
	}

	return fmt.Sprintf("%v,%v", strconv.Itoa(x), strconv.Itoa(y)), nil
}

func readFile() (map[string][]string, int, int) {
	out := make(map[string][]string)
	freq := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// fmt.Println(freq)
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)

	if err != nil {
		panic(err)
	}

	defer fi.Close()

	os := bufio.NewScanner(fi)
	x_max := 0
	y_max := 0

	for os.Scan() {
		line := os.Text()
		if len(line) > y_max {
			y_max = len(line)
		}

		for j, char := range strings.Split(line, "") {
			if strings.Contains(freq, char) {
				out[char] = append(out[char], fmt.Sprintf("%v,%v", x_max, j))
			}
		}
		x_max++
	}

	return out, x_max, y_max
}
