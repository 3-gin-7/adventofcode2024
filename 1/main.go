package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right, counter := readFile()

	// fmt.Println(left)
	// fmt.Println(right)

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var total_dif int

	for i := 0; i < len(left); i++ {
		float_dif := float64(left[i] - right[i])
		abs_dif := math.Abs(float_dif)
		total_dif += int(abs_dif)
		// fmt.Printf("dif is: %v\r\n", abs_dif)
	}

	sum_mult := 0
	for _, i := range left {
		mult := counter[i] * i
		sum_mult += mult
	}

	fmt.Printf("addition output: %v\r\n", total_dif)
	fmt.Printf("multiplication output: %v\r\n", sum_mult)
}

func readFile() ([]int, []int, map[int]int) {
	var left []int
	var right []int
	var counter map[int]int = make(map[int]int)

	fi, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open input.txt")
	}

	sc := bufio.NewScanner(fi)

	for sc.Scan() {
		line := sc.Text()
		nums := strings.Split(line, "   ")
		if i, err := strconv.Atoi(nums[0]); err == nil {
			left = append(left, i)
		}

		if i, err := strconv.Atoi(nums[1]); err == nil {
			right = append(right, i)
			counter[i]++
		}
	}

	return left, right, counter
}
