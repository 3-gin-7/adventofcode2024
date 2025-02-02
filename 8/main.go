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
	antenas := make(map[string]int)
	coords, x_max, y_max := readFile()

	// loop over the corrds
	for _, coord_list := range coords {
		// check the coord list length
		if len(coord_list) < 2 {
			continue
		}

		for x := 0; x < len(coord_list); x++ {
			for y := x + 1; y < len(coord_list); y++ {
				first_diff, err1 := getDiff(coord_list[x], coord_list[y], x_max, y_max)
				second_diff, err2 := getDiff(coord_list[y], coord_list[x], x_max, y_max)

				if err1 == nil {
					// add the coords to the listA
					first_coords := strings.Join(first_diff, ",")
					antenas[first_coords]++
				}

				if err2 == nil {
					// add the coords to the listA
					second_coors := strings.Join(second_diff, ",")
					antenas[second_coors]++
				}
			}
		}
	}

	fmt.Printf("len is: %v\r\n", len(antenas))
}

func getDiff(s1, s2 string, x_max, y_max int) ([]string, error) {
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

	x1 = x1 + x_diff
	y1 = y1 + y_diff

	if x1 < 0 || y1 < 0 {
		return []string{}, errors.New("invalid coords")
	}

	if x1 >= x_max || y1 >= y_max {
		return []string{}, errors.New("invalid coords")
	}

	return []string{strconv.Itoa(x1), strconv.Itoa(y1)}, nil
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
