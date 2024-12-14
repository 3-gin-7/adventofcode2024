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
		fmt.Println(i)
		replica1 := readSlice(i)
		replica2 := readSlice(i)

		isValid, index := validate(i)

		if !isValid {
			tmpRetry := false
			if index != 0 {
				damp1 := append(replica1[:index-1], replica1[index:]...)
				retry1, _ := validate(damp1)
				tmpRetry = retry1
			}
			damp2 := append(replica2[:index], replica2[index+1:]...)
			damp3 := append(i[:index+1], i[index+2:]...)
			retry2, _ := validate(damp2)
			retry3, _ := validate(damp3)

			check := tmpRetry || retry2 || retry3

			if check {
				count++
			}
		} else {
			count++
		}
	}

	fmt.Printf("output is: %v", count)
}

// determine if the sequence is *correct and return first index where it is not correct
// *correct -  means that the sequence is either increasing or decreasing by no more than 3
func validate(i []int) (bool, int) {
	isValid := false
	index := 0

	// determine if sequence is increasing or decreasing
	if (i[0] - i[1]) == 0 {
		// seq is not increasing
		return isValid, 0
	}

	seq_inc := (i[0] - i[1]) < 0

	for j := 0; j < len(i); j++ {
		index = j
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

	return isValid, index
}

// for some reason go modifies the original array during append
// create new instance of the array
func readSlice(slice []int) []int {
	output := make([]int, len(slice))
	for c, i := range slice {
		output[c] = i
	}

	return output
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
