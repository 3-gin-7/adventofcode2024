package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var registerA int
var registerB int
var registerC int
var pointer int
var program []string
var output []int

func main() {
	// load the registers and program from file
	readFile()
	pointer = 0

	fmt.Printf("A: %v, B: %v, C: %v, Program: %v\n", registerA, registerB, registerC, program)

	for i := 0; i < len(program); i++ {

		if pointer+1 > len(program) {
			fmt.Printf("can not access pointer of: %v\r\n", pointer+1)
			return
		}

		opcode := program[pointer]
		operand, err := strconv.Atoi(program[pointer+1])

		pointer += 2

		if err != nil {
			panic(err)
		}

		op := Operand{opcode, operand}
		op.Process()

		i = pointer
		// fmt.Printf("A: %v, B: %v, C: %v, Program: %v\n", registerA, registerB, registerC, program)
	}

	fmt.Printf("A: %v, B: %v, C: %v, Program: %v\n", registerA, registerB, registerC, program)
	var stringSlice []string
	for _, num := range output {
		stringSlice = append(stringSlice, strconv.Itoa(num))
	}
	fmt.Printf("output: %v\n", strings.Join(stringSlice, ","))

}

func readFile() {
	// FILE_NAME := "test.txt"
	FILE_NAME := "input.txt"

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	sc := bufio.NewScanner(fi)

	for sc.Scan() {
		line := sc.Text()

		// skip if empty
		if line == "" {
			continue
		} else if strings.Contains(strings.ToLower(line), "register") {
			split := strings.Split(line, " ")
			register_name := split[1]
			register_val, err := strconv.Atoi(split[2])

			if err != nil {
				panic(err)
			}

			if strings.Contains(register_name, "A:") {
				registerA = register_val
			} else if strings.Contains(register_name, "B:") {
				registerB = register_val
			} else if strings.Contains(register_name, "C:") {
				registerC = register_val
			}
		} else if strings.Contains(strings.ToLower(line), "program") {
			program_split := strings.Split(line, " ")
			program_val := program_split[1]
			for _, val := range strings.Split(program_val, ",") {
				program = append(program, val)
			}
		}

	}

}

func MathPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}

func (op *Operand) Process() {
	switch op.id {
	case "0":
		{
			numerator := registerA
			power := op.ComboOperand()
			denominator := MathPow(2, power)
			registerA = numerator / denominator
			fmt.Printf("Processing opcode 0 with input: %v\r\n", op.input)
		}
	case "1":
		{
			registerB = registerB ^ op.input
			fmt.Printf("Processing opcode 1 with input: %v\r\n", op.input)
		}
	case "2":
		{
			combo := op.ComboOperand()
			registerB = combo % 8
			fmt.Printf("Processing opcode 2 with input: %v\r\n", op.input)
		}
	case "3":
		{
			if registerA == 0 {
				return
			}
			pointer = op.input
			fmt.Printf("Processing opcode 3 with input: %v\r\n", op.input)
		}
	case "4":
		{
			registerB = registerB ^ registerC
			fmt.Printf("Processing opcode 4 with input: %v\r\n", op.input)
		}
	case "5":
		{
			output_val := op.ComboOperand() % 8
			output = append(output, output_val)
			fmt.Printf("Processing opcode 5 with input: %v\r\n", op.input)
		}
	case "6":
		{
			numerator := registerA
			power := op.ComboOperand()
			denominator := MathPow(2, power)
			registerB = numerator / denominator
			fmt.Printf("Processing opcode 6 with input: %v\r\n", op.input)
		}
	case "7":
		{
			numerator := registerA
			power := op.ComboOperand()
			denominator := MathPow(2, power)
			registerC = numerator / denominator
			fmt.Printf("Processing opcode 7 with input: %v\r\n", op.input)
		}
	default:
		fmt.Println("default")
	}
}

func (op *Operand) ComboOperand() int {
	switch op.input {
	case 0:
		{
			return 0
		}
	case 1:
		{
			return 1
		}
	case 2:
		{
			return 2
		}
	case 3:
		{
			return 3
		}
	case 4:
		{
			return registerA
		}
	case 5:
		{
			return registerB
		}
	case 6:
		{
			return registerC
		}
	default:
		panic("Can not return 7")
	}
}

type Operand struct {
	id    string
	input int
}
