package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	part_one_claws, part_two_claws := readFile()

	final_part_one_sum := 0
	for _, claw := range part_one_claws {
		// Try all possible combinations of button presses up to 100
		final_part_one_sum += partOne(claw)

	}
	fmt.Printf("Final sum for part one is: %d\n", final_part_one_sum)

	final_part_two_sum := 0
	test_sum := 0
	for _, claw := range part_two_claws {
		q := partTwo(claw)
		w := goNumTest(claw)
		if q != w {
			fmt.Println("hey")
			r := goNumTest(claw)
			fmt.Println(r)
		}
		final_part_two_sum += q
		test_sum += w
	}
	fmt.Printf("Final sum for part two is: %d\n", final_part_two_sum)
	fmt.Printf("Test sum for part two is:  %d\n", test_sum)
}

func partOne(claw *ClawInfo) int {

	min_score := 999999
	found_solution := false

	// Try both directions (X->Y and Y->X)
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			// Check X->Y direction
			if a*claw.XInput[0]+b*claw.XInput[1] == claw.XTotal {
				// Verify this solution works for Y total
				if a*claw.YInput[0]+b*claw.YInput[1] == claw.YTotal {
					score := a*3 + b*1
					if score < min_score {
						min_score = score
						found_solution = true
					}
				}
			}

			// Check Y->X direction
			if a*claw.YInput[0]+b*claw.YInput[1] == claw.YTotal {
				// Verify this solution works for X total
				if a*claw.XInput[0]+b*claw.XInput[1] == claw.XTotal {
					score := a*3 + b*1
					if score < min_score {
						min_score = score
						found_solution = true
					}
				}
			}
		}
	}

	if found_solution {
		return min_score
	} else {
		return 0
	}
}

func goNumTest(claw *ClawInfo) int {

	a, b := float64(0), float64(0)

	A := mat.NewDense(2, 2, []float64{
		float64(claw.XInput[0]), float64(claw.XInput[1]),
		float64(claw.YInput[0]), float64(claw.YInput[1]),
	})
	B := mat.NewVecDense(2, []float64{float64(claw.XTotal), float64(claw.YTotal)})
	var x mat.VecDense

	err := x.SolveVec(A, B)
	if err != nil {
		fmt.Println("No exact solution:", err)
		return 0
	}

	a = x.AtVec(0)
	b = x.AtVec(1)

	for i := 0; i < x.Len(); i++ {
		val := x.AtVec(i)
		remainder := val - math.Round(val)
		if math.Abs(remainder) != 0 {
			decimal_val := fmt.Sprintf("%f", x.AtVec(i))
			if !strings.Contains(decimal_val, ".999") && !strings.Contains(decimal_val, ".000") {
				return 0
			} else if strings.Contains(decimal_val, ".999") || strings.Contains(decimal_val, ".000") {
				a = math.Round(x.AtVec(0))
				b = math.Round(x.AtVec(1))
			}
		}
	}

	return int(a*3 + b)
}

func partTwo(claw *ClawInfo) int {
	aX := claw.XInput[0]
	aY := claw.YInput[0]
	bX := claw.XInput[1]
	bY := claw.YInput[1]
	pX := claw.XTotal
	pY := claw.YTotal

	detA := int64(aX)*int64(bY) - int64(aY)*int64(bX)
	if detA == 0 {
		return 0
	}

	detA1 := int64(pX)*int64(bY) - int64(pY)*int64(bX)
	detA2 := int64(aX)*int64(pY) - int64(aY)*int64(pX)

	if detA1%detA != 0 || detA2%detA != 0 {
		return 0
	}

	nA := int(detA1 / detA)
	nB := int(detA2 / detA)

	if nA < 0 || nB < 0 {
		return 0
	}

	cost := 3*nA + nB
	return cost
}

func readFile() ([]*ClawInfo, []*ClawInfo) {
	part_one_output := []*ClawInfo{}
	part_two_output := []*ClawInfo{}
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)
	part_one_claw := &ClawInfo{}
	part_two_claw := &ClawInfo{}

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			part_one_output = append(part_one_output, part_one_claw)
			part_two_output = append(part_two_output, part_two_claw)
			part_one_claw = &ClawInfo{}
			part_two_claw = &ClawInfo{}
			continue
		}

		if strings.Contains(line, "Button") {
			re := regexp.MustCompile(`[0-9]+`)
			matches := re.FindAllString(line, -1)
			if len(matches) > 0 {
				x_number, _ := strconv.Atoi(matches[0])
				y_number, _ := strconv.Atoi(matches[1])

				part_one_claw.XInput = append(part_one_claw.XInput, x_number)
				part_one_claw.YInput = append(part_one_claw.YInput, y_number)

				x_number, _ = strconv.Atoi(matches[0])
				y_number, _ = strconv.Atoi(matches[1])

				part_two_claw.XInput = append(part_two_claw.XInput, x_number)
				part_two_claw.YInput = append(part_two_claw.YInput, y_number)
			}
		} else if strings.Contains(line, "Prize") {
			re := regexp.MustCompile(`[0-9]+`)
			matches := re.FindAllString(line, -1)
			if len(matches) > 0 {
				x_number, _ := strconv.Atoi(matches[0])
				y_number, _ := strconv.Atoi(matches[1])

				part_one_claw.XTotal = x_number
				part_one_claw.YTotal = y_number

				x_number, _ = strconv.Atoi(matches[0])
				y_number, _ = strconv.Atoi(matches[1])

				x_number += 10000000000000
				y_number += 10000000000000
				part_two_claw.XTotal = x_number
				part_two_claw.YTotal = y_number
			}
		}
	}

	// Don't forget to append the last claw if the file doesn't end with a blank line
	if part_one_claw.XInput != nil {
		part_one_output = append(part_one_output, part_one_claw)
		part_two_output = append(part_two_output, part_two_claw)
	}

	return part_one_output, part_two_output
}

type ClawInfo struct {
	XInput []int
	YInput []int
	XTotal int
	YTotal int
}
