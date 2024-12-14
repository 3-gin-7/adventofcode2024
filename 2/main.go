package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 549 too high

func main() {
	input := readFile()

	count := 0
	// loop over outer
	for _, i := range input {

		// special case is 2 and 3??
		// get the seq increment
		// run the validate on seq with increment

		isValid := validate(i)
		// fmt.Println(isValid)

		if isValid {
			count++
		}
		// break
	}

	fmt.Printf("output is: %v", count)
}

func validate(i []int) bool {
	isValid := false

	fmt.Printf("got seq: %v\n\r", i)
	// determine if sequence is increasing or decreasing
	if (i[0] - i[1]) == 0 {
		// seq is not increasing
		return isValid
	}

	seq_inc := (i[0] - i[1]) < 0
	// fmt.Printf("seq inc is: %v\r\n", seq_inc)

	for j := 0; j < len(i); j++ {
		if j == len(i)-1 {
			break
		}

		diff := i[j] - i[j+1]

		if diff == 0 || math.Abs(float64(diff)) > 3 || (diff > 0) == seq_inc {
			isValid = false
			break
		} else {
			isValid = true
		}
	}

	return isValid
}

func readFile() [][]int {
	var output [][]int

	fi, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open input.txt")
	}

	sc := bufio.NewScanner(fi)

	line_counter := 0
	for sc.Scan() {
		tmp := []int{}
		line := sc.Text()
		nums := strings.Split(line, " ")

		// convert strings to ints
		for _, i := range nums {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic("corrupted file")
			}
			tmp = append(tmp, j)
		}

		output = append(output, tmp)

		line_counter++
	}

	return output
}
