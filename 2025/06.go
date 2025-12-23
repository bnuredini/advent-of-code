package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/06.in")
	if err != nil {
		log.Fatalf("failed to read puzzle input: %v", err)
	}
	defer f.Close()

	answer1 := 0
	answer2 := 0

	mat1 := [][]int{}
	lines := []string{}
	colsWithAddition := make(map[int]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)

		row := []int{}
		colIdx := 0

		parts := strings.Split(line, " ")
		for _, part := range parts {
			if strings.TrimSpace(part) == "" {
				continue
			}

			if part == "+" || part == "*" {
				colsWithAddition[colIdx] = (part == "+")
				colIdx += 1
			} else {
				num, _ := strconv.Atoi(part)
				row = append(row, num)
			}
		}

		if len(row) != 0 {
			mat1 = append(mat1, row)
		}
	}

	for colIdx := range len(mat1[0]) {
		colValue := 0
		usingAddition := colsWithAddition[colIdx]

		if !usingAddition {
			colValue = 1
		}

		for rowIdx := range len(mat1) {
			if usingAddition {
				colValue += mat1[rowIdx][colIdx]
			} else {
				colValue *= mat1[rowIdx][colIdx]
			}
		}

		answer1 += colValue
	}

	mat2 := [][]int{}
	numLineCols := len(lines[0])
	numLineRows := len(lines)

	row := []int{}

	for colIdx := range numLineCols {
		lastColCell := lines[numLineRows-1][colIdx]

		if colIdx != 0 && (lastColCell == '+' || lastColCell == '*') {
			mat2 = append(mat2, row)
			row = []int{}
		}

		concatenatedNumber := ""

		for rowIdx := range numLineRows {
			ch := lines[rowIdx][colIdx]
			if ch == ' ' || ch == '+' || ch == '*' {
				continue
			}

			concatenatedNumber += string(ch)
		}

		if strings.TrimSpace(concatenatedNumber) != "" {
			parsedNumber, _ := strconv.Atoi(concatenatedNumber)
			row = append(row, parsedNumber)
		}
	}

	mat2 = append(mat2, row)

	for i, nums := range mat2 {
		value := 0
		usingAddition := colsWithAddition[i]

		if !usingAddition {
			value = 1
		}

		for _, num := range nums {
			if usingAddition {
				value += num
			} else {
				value *= num
			}
		}

		answer2 += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	fmt.Println(answer1)
	fmt.Println(answer2)
}
