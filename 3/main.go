package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	total := 0
	total_commands := 0
	mul_rg_str := `mul\(\d+,\d+\)`
	params_rg_str := `\d+,\d+`
	do_rg_str := `do\(\)`
	dont_rg_str := `don't\(\)`

	search := readFile("input.txt")

	r, _ := regexp.Compile(mul_rg_str)
	r2, _ := regexp.Compile(params_rg_str)
	r3, _ := regexp.Compile(do_rg_str)
	r4, _ := regexp.Compile(dont_rg_str)

	muls := r.FindAllString(search, -1)
	muls_str_indexes := r.FindAllStringIndex(search, -1)
	muls_idx := extractFirstIndex(muls_str_indexes)
	dos_idx := extractFirstIndex(r3.FindAllStringIndex(search, -1))
	donts_idx := extractFirstIndex(r4.FindAllStringIndex(search, -1))

	// fmt.Printf("idx: %v\r\n", muls_idx)
	// fmt.Printf("do idx: %v\r\n", dos_idx)
	// fmt.Printf("dont idx: %v\r\n", donts_idx)

	dont_index := 0
	allowed_index := make([][]int, 1)

	if dos_idx[0] < donts_idx[0] {
		allowed_index = append(allowed_index, []int{0, dos_idx[0]})
	} else {
		allowed_index = append(allowed_index, []int{0, donts_idx[0]})
	}

	for _, i := range dos_idx {
		if dont_index != len(donts_idx)-1 {
			for j := 0; j < len(donts_idx); j++ {
				if donts_idx[j] > i {
					dont_index = j
					break
				}
			}
		} else {
			break
		}
		allowed_index = append(allowed_index, []int{i, donts_idx[dont_index]})
	}

	if dos_idx[len(dos_idx)-1] > donts_idx[len(donts_idx)-1] {
		t := dos_idx[len(dos_idx)-1]
		allowed_index = append(allowed_index, []int{t, muls_idx[len(muls)-1]})
	}

	valid_muls := make([]string, 1)

	muls_index := 0
	for _, i := range allowed_index {
		if len(i) == 0 {
			continue
		}
		lower := i[0]
		upper := i[1]

		if muls_index >= len(muls_str_indexes) {
			break
		}

		for {
			val := muls_idx[muls_index]
			if val >= lower && val <= upper {
				valid_muls = append(valid_muls, muls[muls_index])
			} else if val > upper {
				break
			}
			muls_index++

			if muls_index == len(muls_idx) {
				break
			}
		}
	}

	for _, i := range muls {
		parsed := r2.FindString(i)

		total += getMulResult(parsed)
	}

	fmt.Printf("part one total: %v\r\n", total)

	for index, i := range valid_muls {
		if index == 0 {
			continue
		}
		parsed := r2.FindString(i)

		total_commands += getMulResult(parsed)
	}

	fmt.Printf("part 2 total: %v", total_commands)

}

func getMulResult(s string) int {
	params := strings.Split(s, ",")

	num1, _ := strconv.Atoi(params[0])
	num2, _ := strconv.Atoi(params[1])

	return num1 * num2
}

func extractFirstIndex(idx [][]int) []int {
	output := make([]int, 0)
	for _, i := range idx {
		output = append(output, i[0])
	}

	return output
}

func readFile(filename string) string {

	fi, err := os.Open(filename)

	if err != nil {
		defer fi.Close()
	}

	if err != nil {
		fmt.Printf("failed to open %v\r\n", filename)
	}

	sc := bufio.NewScanner(fi)
	var sb strings.Builder

	for sc.Scan() {
		line := sc.Text()
		sb.WriteString(line)
	}

	return sb.String()
}
