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
	mul_rg_str := `mul\(\d+,\d+\)`
	params_rg_str := `\d+,\d+`
	do_rg_str := `do\(\)`
	dont_rg_str := `don't\(\)`

	// search := readFile("input.txt")
	search := readFile("test2.txt")
	// fmt.Println(search)

	r, _ := regexp.Compile(mul_rg_str)
	r2, _ := regexp.Compile(params_rg_str)
	r3, _ := regexp.Compile(do_rg_str)
	r4, _ := regexp.Compile(dont_rg_str)

	fmt.Println(r3.FindString(""))
	fmt.Println(r4.FindString(""))

	muls := r.FindAllString(search, -1)
	muls_idx := r.FindAllStringIndex(search, -1)
	dos_idx := r3.FindAllStringIndex(search, -1)
	donts_idx := r4.FindAllStringIndex(search, -1)

	fmt.Printf("strings: %v\r\n", muls)
	fmt.Printf("idx: %v\r\n", muls_idx)
	fmt.Printf("do idx: %v\r\n", dos_idx)
	fmt.Printf("dont idx: %v\r\n", donts_idx)

	// get the first index from muls, dos and donts

	for _, i := range muls {
		parsed := r2.FindString(i)

		total += getMulResult(parsed)
	}

	fmt.Printf("total: %v", total)
}

func getMulResult(s string) int {
	params := strings.Split(s, ",")

	num1, _ := strconv.Atoi(params[0])
	num2, _ := strconv.Atoi(params[1])

	return num1 * num2
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
