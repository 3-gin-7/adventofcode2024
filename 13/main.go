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
	claws := readFile()

	final_sum := 0
	for _, claw := range claws {
		// Try all possible combinations of button presses up to 100
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
			final_sum += min_score
		}
	}

	fmt.Printf("Final sum is: %d\n", final_sum)
}

func readFile() []*ClawInfo {
	output := []*ClawInfo{}
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)
	claw := &ClawInfo{}

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			output = append(output, claw)
			claw = &ClawInfo{}
			continue
		}

		if strings.Contains(line, "Button") {
			re := regexp.MustCompile(`[0-9]+`)
			matches := re.FindAllString(line, -1)
			if len(matches) > 0 {
				x_number, _ := strconv.Atoi(matches[0])
				y_number, _ := strconv.Atoi(matches[1])

				claw.XInput = append(claw.XInput, x_number)
				claw.YInput = append(claw.YInput, y_number)
			}
		} else if strings.Contains(line, "Prize") {
			re := regexp.MustCompile(`[0-9]+`)
			matches := re.FindAllString(line, -1)
			if len(matches) > 0 {
				x_number, _ := strconv.Atoi(matches[0])
				y_number, _ := strconv.Atoi(matches[1])

				claw.XTotal = x_number
				claw.YTotal = y_number
			}
		}
	}

	// Don't forget to append the last claw if the file doesn't end with a blank line
	if claw.XInput != nil {
		output = append(output, claw)
	}

	return output
}

type ClawInfo struct {
	XInput []int
	YInput []int
	XTotal int
	YTotal int
}
